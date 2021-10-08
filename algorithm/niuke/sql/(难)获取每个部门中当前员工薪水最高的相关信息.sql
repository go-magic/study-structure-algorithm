drop table if exists  `dept_emp` ; 
drop table if exists  `salaries` ; 
CREATE TABLE `dept_emp` (
`emp_no` int(11) NOT NULL,
`dept_no` char(4) NOT NULL,
`from_date` date NOT NULL,
`to_date` date NOT NULL,
PRIMARY KEY (`emp_no`,`dept_no`));
CREATE TABLE `salaries` (
`emp_no` int(11) NOT NULL,
`salary` int(11) NOT NULL,
`from_date` date NOT NULL,
`to_date` date NOT NULL,
PRIMARY KEY (`emp_no`,`from_date`));
INSERT INTO dept_emp VALUES(10001,'d001','1986-06-26','9999-01-01');
INSERT INTO dept_emp VALUES(10002,'d001','1996-08-03','9999-01-01');

INSERT INTO salaries VALUES(10001,88958,'2002-06-22','9999-01-01');
INSERT INTO salaries VALUES(10002,72527,'2001-08-02','9999-01-01');

-- 解题思路: 先把两张表合成一张表，再根据order by 查询出每个部门的最高工资
SELECT
	d.dept_no,d.emp_no,MAX(s.salary) maxSalary
FROM
	dept_emp d
left JOIN
	salaries s
on
	s.emp_no=d.emp_no
GROUP BY
	d.dept_no
ORDER BY
	d.dept_no
