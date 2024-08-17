INSERT INTO loan(id, borrower_id, amount, rate, approval_proof_url, agreement_letter_url, status, created_by, approved_by, disbursed_by, created_at, updated_at, approved_at, invested_at, disbursed_at)
VALUES
(1, 1, 1000000, 10, '', '', 1, 1, 0, 0, NOW(), NOW(), NULL, NULL, NULL),
(2, 2, 3000000, 12, '', '', 1, 1, 0, 0, NOW(), NOW(), NULL, NULL, NULL),
(3, 3, 350000000, 5, '', '', 1, 2, 0, 0, NOW(), NOW(), NULL, NULL, NULL),
(4, 3, 350000000, 8, 'http://127.0.0.1/upload/proof_3.jpeg', '', 2, 2, 2, 0, NOW(), NOW(), NOW(), NULL, NULL);

INSERT INTO investment(id, investor_id, loan_id, amount, roi, status, created_at, updated_at)
VALUES
(1, 1, 4, 15000000, 1200000, 1, NOW(), NOW()),
(2, 2, 4, 10000000, 800000, 1, NOW(), NOW());

INSERT INTO employee(id, name, status, created_at, updated_at)
VALUES
(1, 'Employee A', 1, NOW(), NOW()),
(2, 'Employee B', 1, NOW(), NOW());

INSERT INTO borrower(id, identification_number, name, status, created_at, updated_at)
VALUES
(1, '1234567890123451', 'Borrower A', 1, NOW(), NOW()),
(2, '1234567890123452', 'Borrower B', 1, NOW(), NOW()),
(3, '1234567890123453', 'Borrower B', 1, NOW(), NOW());

INSERT INTO investor(id, identification_number, name, email, status, created_at, updated_at)
VALUES
(1, '1234567890123454', 'Investor A', 'evinvilan@gmail.com', 1, NOW(), NOW()),
(2, '1234567890123455', 'Investor B', 'cintiawan.evin@gmail.com',  1, NOW(), NOW());