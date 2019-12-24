select
  MAX(Salary) as SecondHighestSalary
from
  Employee
where
  Salary != (select
    MAX(Salary)
  from
    Employee
  )
;
