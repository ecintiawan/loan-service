package action

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/ecintiawan/loan-service/internal/constant"
	"github.com/ecintiawan/loan-service/internal/entity"
	"github.com/ecintiawan/loan-service/internal/repository"
	"github.com/ecintiawan/loan-service/internal/service"
	"github.com/ecintiawan/loan-service/pkg/config"
	"github.com/ecintiawan/loan-service/pkg/file"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewLoanActionImpl(t *testing.T) {
	type args struct {
		config         *config.Config
		repoLoan       repository.Loan
		repoInvestment repository.Investment
		repoInvestor   repository.Investor
		repoUpload     repository.Upload
		repoNotifier   repository.Notifier
		pdfGenerator   file.PDFGenerator
	}
	tests := []struct {
		name string
		args args
		want service.LoanAction
	}{
		{
			name: "success",
			args: args{},
			want: &LoanActionImpl{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLoanActionImpl(tt.args.config, tt.args.repoLoan, tt.args.repoInvestment, tt.args.repoInvestor, tt.args.repoUpload, tt.args.repoNotifier, tt.args.pdfGenerator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLoanActionImpl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoanActionImpl_Approve(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		config         *config.Config
		repoLoan       repository.Loan
		repoInvestment repository.Investment
		repoInvestor   repository.Investor
		repoUpload     repository.Upload
		repoNotifier   repository.Notifier
		pdfGenerator   file.PDFGenerator
	}
	type args struct {
		ctx context.Context
		req *entity.LoanProceed
	}
	defaultArgs := args{
		ctx: context.Background(),
		req: &entity.LoanProceed{
			ApprovalProof: entity.File{
				File:    []byte{123},
				FileExt: ".png",
			},
			Data: &entity.Loan{
				ID:         3,
				ApprovedBy: 2,
				ApprovedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repoUpload: func() *repository.MockUpload {
					mock := repository.NewMockUpload(ctrl)
					mock.EXPECT().
						Upload(gomock.Any(), &entity.File{
							File:     []byte{123},
							FileExt:  ".png",
							FileName: "approval_proof_3.png",
						}).
						Return("http://127.0.0.1:8080/approval_proof_3.png", nil)

					return mock
				}(),
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Update(gomock.Any(), &entity.Loan{
							ID:               3,
							ApprovalProofURL: "http://127.0.0.1:8080/approval_proof_3.png",
							Status:           constant.StatusApproved,
							ApprovedBy:       2,
							ApprovedAt:       time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						}).
						Return(nil)

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name:   "invalid approved at",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					Data: &entity.Loan{},
				},
			},
			wantErr: true,
		},
		{
			name:   "invalid approved by",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					Data: &entity.Loan{
						ApprovedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
					},
				},
			},
			wantErr: true,
		},
		{
			name:   "invalid file content",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					ApprovalProof: entity.File{},
					Data: &entity.Loan{
						ApprovedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						ApprovedBy: 2,
					},
				},
			},
			wantErr: true,
		},
		{
			name:   "invalid file extension",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					ApprovalProof: entity.File{
						File:    []byte{123},
						FileExt: ".pdf",
					},
					Data: &entity.Loan{
						ApprovedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						ApprovedBy: 2,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error on upload",
			fields: fields{
				repoUpload: func() *repository.MockUpload {
					mock := repository.NewMockUpload(ctrl)
					mock.EXPECT().
						Upload(gomock.Any(), &entity.File{
							File:     []byte{123},
							FileExt:  ".png",
							FileName: "approval_proof_3.png",
						}).
						Return("", assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on update",
			fields: fields{
				repoUpload: func() *repository.MockUpload {
					mock := repository.NewMockUpload(ctrl)
					mock.EXPECT().
						Upload(gomock.Any(), &entity.File{
							File:     []byte{123},
							FileExt:  ".png",
							FileName: "approval_proof_3.png",
						}).
						Return("http://127.0.0.1:8080/approval_proof_3.png", nil)

					return mock
				}(),
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Update(gomock.Any(), &entity.Loan{
							ID:               3,
							ApprovalProofURL: "http://127.0.0.1:8080/approval_proof_3.png",
							Status:           constant.StatusApproved,
							ApprovedBy:       2,
							ApprovedAt:       time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						}).
						Return(assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LoanActionImpl{
				config:         tt.fields.config,
				repoLoan:       tt.fields.repoLoan,
				repoInvestment: tt.fields.repoInvestment,
				repoInvestor:   tt.fields.repoInvestor,
				repoUpload:     tt.fields.repoUpload,
				repoNotifier:   tt.fields.repoNotifier,
				pdfGenerator:   tt.fields.pdfGenerator,
			}
			if err := a.Approve(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("LoanActionImpl.Approve() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoanActionImpl_Invest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		config         *config.Config
		repoLoan       repository.Loan
		repoInvestment repository.Investment
		repoInvestor   repository.Investor
		repoUpload     repository.Upload
		repoNotifier   repository.Notifier
		pdfGenerator   file.PDFGenerator
	}
	type args struct {
		ctx context.Context
		req *entity.LoanProceed
	}
	defaultArgs := args{
		ctx: context.Background(),
		req: &entity.LoanProceed{
			Data: &entity.Loan{
				ID:         3,
				Amount:     2000000,
				InvestedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				config: &config.Config{},
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Update(gomock.Any(), &entity.Loan{
							ID:         3,
							Amount:     2000000,
							Status:     constant.StatusInvested,
							InvestedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						}).
						Return(nil)

					return mock
				}(),
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), &entity.InvestmentFilter{
							LoanID: 3,
							Status: constant.GeneralStatusActive,
						}).
						Return(float64(2000000), nil)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.InvestmentResult{
							List: []*entity.Investment{
								{
									ID:     1,
									Amount: 2000000,
								},
							},
						}, nil)

					return mock
				}(),
				repoInvestor: func() *repository.MockInvestor {
					mock := repository.NewMockInvestor(ctrl)
					mock.EXPECT().
						GetDetail(gomock.Any(), gomock.Any()).
						Return(&entity.Investor{
							ID:    1,
							Name:  "ole",
							Email: "test@gmail.com",
						}, nil).AnyTimes()

					return mock
				}(),
				repoNotifier: func() *repository.MockNotifier {
					mock := repository.NewMockNotifier(ctrl)
					mock.EXPECT().
						Notify(gomock.Any(), gomock.Any()).
						Return(nil).AnyTimes()

					return mock
				}(),
				pdfGenerator: func() *file.MockPDFGenerator {
					mock := file.NewMockPDFGenerator(ctrl)
					mock.EXPECT().
						Generate(gomock.Any()).
						Return([]byte{123}, nil).AnyTimes()

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name:   "invalid param",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					Data: &entity.Loan{},
				},
			},
			wantErr: true,
		},
		{
			name: "error on get amount sum",
			fields: fields{
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), &entity.InvestmentFilter{
							LoanID: 3,
							Status: constant.GeneralStatusActive,
						}).
						Return(float64(0), assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "invalid investment sum",
			fields: fields{
				config: &config.Config{},
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), &entity.InvestmentFilter{
							LoanID: 3,
							Status: constant.GeneralStatusActive,
						}).
						Return(float64(20000000), nil)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on update loan",
			fields: fields{
				config: &config.Config{},
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Update(gomock.Any(), &entity.Loan{
							ID:         3,
							Amount:     2000000,
							Status:     constant.StatusInvested,
							InvestedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						}).
						Return(assert.AnError)

					return mock
				}(),
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), &entity.InvestmentFilter{
							LoanID: 3,
							Status: constant.GeneralStatusActive,
						}).
						Return(float64(2000000), nil)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on getting investment",
			fields: fields{
				config: &config.Config{},
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Update(gomock.Any(), &entity.Loan{
							ID:         3,
							Amount:     2000000,
							Status:     constant.StatusInvested,
							InvestedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						}).
						Return(nil)

					return mock
				}(),
				repoInvestment: func() *repository.MockInvestment {
					mock := repository.NewMockInvestment(ctrl)
					mock.EXPECT().
						GetAmountSum(gomock.Any(), &entity.InvestmentFilter{
							LoanID: 3,
							Status: constant.GeneralStatusActive,
						}).
						Return(float64(2000000), nil)
					mock.EXPECT().
						Get(gomock.Any(), gomock.Any()).
						Return(entity.InvestmentResult{}, assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LoanActionImpl{
				config:         tt.fields.config,
				repoLoan:       tt.fields.repoLoan,
				repoInvestment: tt.fields.repoInvestment,
				repoInvestor:   tt.fields.repoInvestor,
				repoUpload:     tt.fields.repoUpload,
				repoNotifier:   tt.fields.repoNotifier,
				pdfGenerator:   tt.fields.pdfGenerator,
			}
			if err := a.Invest(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("LoanActionImpl.Invest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoanActionImpl_Disburse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		config         *config.Config
		repoLoan       repository.Loan
		repoInvestment repository.Investment
		repoInvestor   repository.Investor
		repoUpload     repository.Upload
		repoNotifier   repository.Notifier
		pdfGenerator   file.PDFGenerator
	}
	type args struct {
		ctx context.Context
		req *entity.LoanProceed
	}
	defaultArgs := args{
		ctx: context.Background(),
		req: &entity.LoanProceed{
			AgreementLetter: entity.File{
				File:    []byte{123},
				FileExt: ".pdf",
			},
			Data: &entity.Loan{
				ID:          3,
				DisbursedBy: 2,
				DisbursedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				repoUpload: func() *repository.MockUpload {
					mock := repository.NewMockUpload(ctrl)
					mock.EXPECT().
						Upload(gomock.Any(), &entity.File{
							File:     []byte{123},
							FileExt:  ".pdf",
							FileName: "agreement_letter_3.pdf",
						}).
						Return("http://127.0.0.1:8080/agreement_letter_3.pdf", nil)

					return mock
				}(),
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Update(gomock.Any(), &entity.Loan{
							ID:                 3,
							AgreementLetterURL: "http://127.0.0.1:8080/agreement_letter_3.pdf",
							Status:             constant.StatusDisbursed,
							DisbursedBy:        2,
							DisbursedAt:        time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						}).
						Return(nil)

					return mock
				}(),
			},
			args: defaultArgs,
		},
		{
			name:   "invalid disbursed at",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					Data: &entity.Loan{},
				},
			},
			wantErr: true,
		},
		{
			name:   "invalid disbursed by",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					Data: &entity.Loan{
						DisbursedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
					},
				},
			},
			wantErr: true,
		},
		{
			name:   "invalid file content",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					AgreementLetter: entity.File{},
					Data: &entity.Loan{
						DisbursedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						DisbursedBy: 2,
					},
				},
			},
			wantErr: true,
		},
		{
			name:   "invalid file extension",
			fields: fields{},
			args: args{
				ctx: context.Background(),
				req: &entity.LoanProceed{
					AgreementLetter: entity.File{
						File:    []byte{123},
						FileExt: ".png",
					},
					Data: &entity.Loan{
						DisbursedAt: time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						DisbursedBy: 2,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "error on upload",
			fields: fields{
				repoUpload: func() *repository.MockUpload {
					mock := repository.NewMockUpload(ctrl)
					mock.EXPECT().
						Upload(gomock.Any(), &entity.File{
							File:     []byte{123},
							FileExt:  ".pdf",
							FileName: "agreement_letter_3.pdf",
						}).
						Return("", assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
		{
			name: "error on update",
			fields: fields{
				repoUpload: func() *repository.MockUpload {
					mock := repository.NewMockUpload(ctrl)
					mock.EXPECT().
						Upload(gomock.Any(), &entity.File{
							File:     []byte{123},
							FileExt:  ".pdf",
							FileName: "agreement_letter_3.pdf",
						}).
						Return("http://127.0.0.1:8080/agreement_letter_3.pdf", nil)

					return mock
				}(),
				repoLoan: func() *repository.MockLoan {
					mock := repository.NewMockLoan(ctrl)
					mock.EXPECT().
						Update(gomock.Any(), &entity.Loan{
							ID:                 3,
							AgreementLetterURL: "http://127.0.0.1:8080/agreement_letter_3.pdf",
							Status:             constant.StatusDisbursed,
							DisbursedBy:        2,
							DisbursedAt:        time.Date(2024, 8, 17, 13, 58, 0, 0, time.Local),
						}).
						Return(assert.AnError)

					return mock
				}(),
			},
			args:    defaultArgs,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &LoanActionImpl{
				config:         tt.fields.config,
				repoLoan:       tt.fields.repoLoan,
				repoInvestment: tt.fields.repoInvestment,
				repoInvestor:   tt.fields.repoInvestor,
				repoUpload:     tt.fields.repoUpload,
				repoNotifier:   tt.fields.repoNotifier,
				pdfGenerator:   tt.fields.pdfGenerator,
			}
			if err := a.Disburse(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("LoanActionImpl.Disburse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
