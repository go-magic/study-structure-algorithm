drop table if exists  `salaries` ; 
CREATE TABLE `salaries` (
`emp_no` int(11) NOT NULL,
`salary` int(11) NOT NULL,
`from_date` date NOT NULL,
`to_date` date NOT NULL,
PRIMARY KEY (`emp_no`,`from_date`));
INSERT INTO salaries VALUES(10001,88958,'2002-06-22','9999-01-01');
INSERT INTO salaries VALUES(10002,72527,'2001-08-02','9999-01-01');
INSERT INTO salaries VALUES(10003,43311,'2001-12-01','9999-01-01');


-- case 1

select
    ss.emp_no,ss.salary
from
    salaries ss
JOIN
    (select 
        distinct(s.salary) sl
    FROM
        salaries s
    order by 
        s.salary desc
    limit
        1,1)
    ssalary
ON
    ss.salary=ssalary.sl
