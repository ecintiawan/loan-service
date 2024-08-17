package action

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/ecintiawan/loan-service/pkg/currency"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/ecintiawan/loan-service/pkg/file"
)

type (
	LoanActionImpl struct {
		config         *config.Config
		repoLoan       repository.Loan
		repoInvestment repository.Investment
		repoInvestor   repository.Investor
		repoUpload     repository.Upload
		repoNotifier   repository.Notifier
		pdfGenerator   file.PDFGenerator
	}
)

func NewLoanActionImpl(
	config *config.Config,
	repoLoan repository.Loan,
	repoInvestment repository.Investment,
	repoInvestor repository.Investor,
	repoUpload repository.Upload,
	repoNotifier repository.Notifier,
	pdfGenerator file.PDFGenerator,
) service.LoanAction {
	return &LoanActionImpl{
		config:         config,
		repoLoan:       repoLoan,
		repoInvestment: repoInvestment,
		repoInvestor:   repoInvestor,
		repoUpload:     repoUpload,
		repoNotifier:   repoNotifier,
		pdfGenerator:   pdfGenerator,
	}
}

func (a *LoanActionImpl) Approve(ctx context.Context, req *entity.LoanProceed) error {
	var (
		err error
	)

	if req.Data.ApprovedAt.IsZero() {
		return errorwrapper.E("invalid approval timestamp", errorwrapper.CodeInvalid)
	}
	if req.Data.ApprovedBy <= 0 {
		return errorwrapper.E("invalid approver ID", errorwrapper.CodeInvalid)
	}

	// process upload
	// simulate upload approval proof file
	// this will upload to our local path instead of a dedicated upload engine
	if req.ApprovalProof.File == nil {
		return errorwrapper.E("invalid approval proof file", errorwrapper.CodeInvalid)
	}
	if !req.IsValidExt() {
		return errorwrapper.E("invalid approval proof file extension", errorwrapper.CodeInvalid)
	}
	req.ApprovalProof.FileName = fmt.Sprintf("approval_proof_%d%s", req.Data.ID, req.ApprovalProof.FileExt)

	req.Data.ApprovalProofURL, err = a.repoUpload.Upload(ctx, &req.ApprovalProof)
	if err != nil {
		return err
	}

	req.Data.Status = constant.StatusApproved

	return a.repoLoan.Update(ctx, req.Data)
}

func (a *LoanActionImpl) Invest(ctx context.Context, req *entity.LoanProceed) error {
	var (
		err error
	)

	if req.Data.InvestedAt.IsZero() {
		return errorwrapper.E("invalid invested timestamp", errorwrapper.CodeInvalid)
	}

	// validate amount sum
	// if investment sum != principal amount, return error
	investmentSum, err := a.repoInvestment.GetAmountSum(ctx, &entity.InvestmentFilter{
		LoanID: req.Data.ID,
		Status: constant.GeneralStatusActive,
	})
	if err != nil {
		return err
	}
	if investmentSum != req.Data.Amount {
		return errorwrapper.E("investment sum doesn't equal the loan's principle amount", errorwrapper.CodeInvalid)
	}

	req.Data.Status = constant.StatusInvested

	err = a.repoLoan.Update(ctx, req.Data)
	if err != nil {
		return err
	}

	// notify all related investors and send them agreement letter
	// first, get all related investors data
	investment, err := a.repoInvestment.Get(ctx, &entity.InvestmentFilter{
		DataTable: entity.DataTableFilter{
			Pagination: entity.DataTablePagination{
				DisablePagination: true,
			},
		},
		LoanID: req.Data.ID,
		Status: constant.GeneralStatusActive,
	})
	if err != nil {
		return err
	}
	a.notifyBulkInvestor(investment.List, req.Data)

	return nil
}

func (a *LoanActionImpl) Disburse(ctx context.Context, req *entity.LoanProceed) error {
	var (
		err error
	)

	if req.Data.DisbursedAt.IsZero() {
		return errorwrapper.E("invalid disbursement timestamp", errorwrapper.CodeInvalid)
	}
	if req.Data.DisbursedBy <= 0 {
		return errorwrapper.E("invalid disburser ID", errorwrapper.CodeInvalid)
	}

	// process upload
	// simulate upload agreement letter file
	// this will upload to our local path instead of a dedicated upload engine
	if req.AgreementLetter.File == nil {
		return errorwrapper.E("invalid agreement letter file", errorwrapper.CodeInvalid)
	}
	if !req.IsValidExt() {
		return errorwrapper.E("invalid agreement letter file extension", errorwrapper.CodeInvalid)
	}
	req.AgreementLetter.FileName = fmt.Sprintf("agreement_letter_%d%s", req.Data.ID, req.AgreementLetter.FileExt)

	req.Data.AgreementLetterURL, err = a.repoUpload.Upload(ctx, &req.AgreementLetter)
	if err != nil {
		return err
	}

	req.Data.Status = constant.StatusDisbursed

	return a.repoLoan.Update(ctx, req.Data)
}

func (a *LoanActionImpl) notifyBulkInvestor(
	investments []*entity.Investment,
	loan *entity.Loan,
) {
	var wg sync.WaitGroup

	for _, val := range investments {
		wg.Add(1)

		go func(investment *entity.Investment) {
			defer wg.Done()

			investor, err := a.repoInvestor.GetDetail(context.TODO(), investment.InvestorID)
			if err != nil {
				log.Println("error getting investor detail", err)
				return
			}

			if err = a.notifyInvestor(investor, investment, loan); err != nil {
				log.Println("error notifying investor", err)
			}
		}(val)
	}
}

func (a *LoanActionImpl) notifyInvestor(
	investor *entity.Investor,
	investment *entity.Investment,
	loan *entity.Loan,
) error {
	finalInvestedAmount := investment.Amount + (investment.Amount * investment.ROI / 100)

	pdfContent := fmt.Sprintf(
		agreementLetterFormat,
		time.Now().Format(constant.DateBeautifyFormat),
		investor.Name,
		investor.Name,
		loan.ID,
		loan.ID,
		currency.ToRupiahFormat(loan.Amount),
		currency.ToRupiahFormat(investment.Amount),
		fmt.Sprintf("%.2f%%", loan.Rate),
		fmt.Sprintf("%.2f%%", investment.ROI),
		currency.ToRupiahFormat(finalInvestedAmount),
		investment.CreatedAt.Format(constant.DateBeautifyFormat),
	)
	fileBytes, err := a.pdfGenerator.Generate(pdfContent)
	if err != nil {
		return err
	}

	emailBodyContent := fmt.Sprintf(agreementEmailFormat,
		investor.Name,
		loan.ID,
		investment.CreatedAt.Format(constant.DateBeautifyFormat),
		currency.ToRupiahFormat(investment.Amount),
		fmt.Sprintf("%.2f%%", investment.ROI),
		currency.ToRupiahFormat(finalInvestedAmount),
	)
	return a.repoNotifier.Notify(context.TODO(), &entity.Notifier{
		To:      []string{investor.Email},
		Subject: fmt.Sprintf("Agreement Letter - Loan ID %d", loan.ID),
		Body:    emailBodyContent,
		Attachment: entity.File{
			File:     fileBytes,
			FileName: fmt.Sprintf(a.config.Vendor.DefaultAgreementLetter.DestFileName, investor.Name),
		},
	})
}
