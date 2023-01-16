package employeestorage

import (
	employeefilter "server/module/employee/filter"
	employeemodel "server/module/employee/model"
)
func (s *sqlStore) GetEmployees(cond *employeefilter.EmployeeFilter) ([]employeemodel.Employee, error){
	strSQL := `
				SELECT
					E.ID,
					E.USERNAME,
					E.FULLNAME,
					E.PHONE,
					R.ROLE
				FROM
					EMPLOYEES E
				INNER JOIN
					ROLES R
				WHERE
					DEL_FLG = 1
	`
	var result []employeemodel.Employee
	s.db.Raw(strSQL, cond.Username, cond.Fullname, cond.Phone).Scan(&result)
	
	return result, nil
}