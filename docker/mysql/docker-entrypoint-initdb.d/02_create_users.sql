CREATE USER IF NOT EXISTS 'sample1_user'@'%' IDENTIFIED BY 'sample1_password';
GRANT ALL ON sample1.* TO 'sample1_user'@'%';

CREATE USER IF NOT EXISTS 'sample2_user'@'%' IDENTIFIED BY 'sample2_password';
GRANT ALL ON sample2.* TO 'sample2_user'@'%';