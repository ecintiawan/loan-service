package handler

import (
	"io"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/pkg/errorwrapper"
	"github.com/labstack/echo/v4"
)

func transformToLoanFilter(c echo.Context) *entity.LoanFilter {
	var (
		filter = &entity.LoanFilter{}
	)

	filter.DataTable.Sort.Field = c.QueryParam("sorted_field")
	filter.DataTable.Sort.Direction = c.QueryParam("sorted_direction")
	filter.DataTable.Pagination.Page, _ = strconv.ParseInt(c.QueryParam("page"), 10, 64)
	filter.DataTable.Pagination.Limit, _ = strconv.ParseInt(c.QueryParam("row"), 10, 64)

	filter.ID, _ = strconv.ParseInt(c.QueryParam("id"), 10, 64)
	filter.BorrowerID, _ = strconv.ParseInt(c.QueryParam("borrower_id"), 10, 64)
	status, _ := strconv.Atoi(c.QueryParam("status"))
	filter.Status = constant.LoanStatus(status)
	filter.CreatedAtStart, _ = time.Parse(constant.TimeISOFormat, c.QueryParam("created_at_start"))
	filter.CreatedAtEnd, _ = time.Parse(constant.TimeISOFormat, c.QueryParam("created_at_end"))
	filter.UpdatedAtStart, _ = time.Parse(constant.TimeISOFormat, c.QueryParam("updated_at_start"))
	filter.UpdatedAtEnd, _ = time.Parse(constant.TimeISOFormat, c.QueryParam("updated_at_end"))
	filter.ApprovedBy, _ = strconv.ParseInt(c.QueryParam("approved_by"), 10, 64)
	filter.DisbursedBy, _ = strconv.ParseInt(c.QueryParam("disbursed_by"), 10, 64)

	return filter
}

func transformToInvestmentFilter(c echo.Context) *entity.InvestmentFilter {
	var (
		filter = &entity.InvestmentFilter{}
	)

	filter.DataTable.Sort.Field = c.QueryParam("sorted_field")
	filter.DataTable.Sort.Direction = c.QueryParam("sorted_direction")
	filter.DataTable.Pagination.Page, _ = strconv.ParseInt(c.QueryParam("page"), 10, 64)
	filter.DataTable.Pagination.Limit, _ = strconv.ParseInt(c.QueryParam("row"), 10, 64)

	filter.ID, _ = strconv.ParseInt(c.QueryParam("id"), 10, 64)
	filter.InvestorID, _ = strconv.ParseInt(c.QueryParam("investor_id"), 10, 64)
	filter.LoanID, _ = strconv.ParseInt(c.QueryParam("loan_id"), 10, 64)
	filter.Status, _ = strconv.Atoi(c.QueryParam("status"))
	filter.CreatedAtStart, _ = time.Parse(constant.TimeISOFormat, c.QueryParam("created_at_start"))
	filter.CreatedAtEnd, _ = time.Parse(constant.TimeISOFormat, c.QueryParam("created_at_end"))
	filter.UpdatedAtStart, _ = time.Parse(constant.TimeISOFormat, c.QueryParam("updated_at_start"))
	filter.UpdatedAtEnd, _ = time.Parse(constant.TimeISOFormat, c.QueryParam("updated_at_end"))

	return filter
}

func transformToLoanProceed(c echo.Context) *entity.LoanProceed {
	var (
		req = &entity.LoanProceed{}
	)

	action, _ := strconv.ParseInt(c.FormValue("action"), 10, 64)
	req.Action = constant.LoanAction(action)
	req.ApprovalProof.File, req.ApprovalProof.FileExt, _ = readMultipartFile(c, "approval_proof")
	req.AgreementLetter.File, req.AgreementLetter.FileExt, _ = readMultipartFile(c, "agreement_letter")
	req.Data = &entity.Loan{}
	req.Data.ID, _ = strconv.ParseInt(c.Param("id"), 10, 64)
	req.Data.ApprovedBy, _ = strconv.ParseInt(c.FormValue("approved_by"), 10, 64)
	req.Data.ApprovedAt, _ = time.Parse(constant.TimeISOFormat, c.FormValue("approved_at"))
	req.Data.DisbursedBy, _ = strconv.ParseInt(c.FormValue("disbursed_by"), 10, 64)
	req.Data.DisbursedAt, _ = time.Parse(constant.TimeISOFormat, c.FormValue("disbursed_at"))
	req.Data.InvestedAt, _ = time.Parse(constant.TimeISOFormat, c.FormValue("invested_at"))

	return req
}

func readMultipartFile(c echo.Context, key string) ([]byte, string, error) {
	file, header, err := c.Request().FormFile(key)
	if err != nil {
		return nil, "", errorwrapper.E("error getting form file", errorwrapper.CodeInvalid)
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)

	bFile, err := io.ReadAll(file)
	if err != nil {
		return nil, "", errorwrapper.E("error reading form file", errorwrapper.CodeInvalid)
	}

	return bFile, ext, nil
}
