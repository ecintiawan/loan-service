package action

const (
	agreementLetterFormat = `Date: %s|Subject: Agreement Letter for Loan Investment|To:|%s||Dear %s,|We are pleased to confirm your investment in the loan offered by Company A to Loan ID %d. Below are the details of the investment:||- Loan ID: %d|- Principal Amount: %s|- Initial Invested Amount: %s|- Interest Rate: %s|- Return on Investment (ROI): %s|- Final Invested Amount: %s|- Investment Date: %s||Enclosed with this letter, please find the signed agreement letter. Kindly review the attached document and retain it for your records.|Should you have any questions or require further information, please do not hesitate to contact us.||Thank you for your trust and investment in Company A.||Best regards,|Admin|Company A`
	agreementEmailFormat  = `Dear %s,

Attached is the agreement letter for your recent investment in Loan %d. Please review and keep this document for your records.

Below is the detail of your investment:
- Investment Date: %s
- Initial Investment Amount: %s
- ROI: %s
- Final Investment Amount: %s

Thank you for your trust in us.`
)
