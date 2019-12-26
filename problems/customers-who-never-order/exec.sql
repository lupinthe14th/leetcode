select
  a.name as Customers
from
  Customers as a
  left join Orders as b on a.id = b.customerid
where
  b.id is null
