select
  b.Name as Employee
from
  Employee as a,
  Employee as b
where
  a.Id = b.ManagerId
  and a.Salary < b.Salary
;
