CREATE
DATABASE loan_service;

USE
loan_service;

CREATE TABLE loans
(
    id                   INT AUTO_INCREMENT PRIMARY KEY,
    borrower_id          VARCHAR(255)   NOT NULL,
    principal_amount     DECIMAL(15, 2) NOT NULL,
    rate                 DECIMAL(5, 2)  NOT NULL,
    roi                  DECIMAL(5, 2)  NOT NULL,
    agreement_letter_url VARCHAR(255)   NOT NULL,
    state                ENUM('proposed', 'approved', 'invested', 'disbursed') NOT NULL DEFAULT 'proposed',
    created_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at           TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE approval_info
(
    id            INT AUTO_INCREMENT PRIMARY KEY,
    loan_id       INT          NOT NULL,
    picture_proof VARCHAR(255) NOT NULL,
    employee_id   VARCHAR(255) NOT NULL,
    date          DATE         NOT NULL,
    FOREIGN KEY (loan_id) REFERENCES loans (id) ON DELETE CASCADE
);

CREATE TABLE investments
(
    id          INT AUTO_INCREMENT PRIMARY KEY,
    loan_id     INT            NOT NULL,
    investor_id VARCHAR(255)   NOT NULL,
    amount      DECIMAL(15, 2) NOT NULL,
    FOREIGN KEY (loan_id) REFERENCES loans (id) ON DELETE CASCADE
);

CREATE TABLE disbursement_info
(
    id                   INT AUTO_INCREMENT PRIMARY KEY,
    loan_id              INT          NOT NULL,
    employee_id          VARCHAR(255) NOT NULL,
    agreement_letter_url VARCHAR(255) NOT NULL,
    date                 DATE         NOT NULL,
    FOREIGN KEY (loan_id) REFERENCES loans (id) ON DELETE CASCADE
);