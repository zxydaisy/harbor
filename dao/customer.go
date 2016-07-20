package dao

import (
	"github.com/vmware/harbor/models"
	"fmt"
	"github.com/caicloud/fornax/pkg/log"
)


func AddCustomer(m models.Customer) (res bool , err error) {
	o := GetOrmer()

	sql := `select * from customer where name= ?`
	type dummy struct{}
	var d []dummy
	_, err = o.Raw(sql, m.Name).QueryRows(&d)
	if len(d) != 0 {
		return false, err
	}

	sql = `insert into customer(name,tag) values(?, ?)`
	p,_ := o.Raw(sql).Prepare()
	defer p.Close()

	_,err = p.Exec(m.Name,m.Tag )

	return true,err
}

func GetCustomerById(id int) (*models.Customer, error) {
	o := GetOrmer()

	p := models.Customer{}
	err := o.Raw("select * from customer where id = ?", id).QueryRow(&p)

	if err != nil {
		return nil, err
	}

	log.Infof("user: %+v", p)

	return &p, nil
}

func GetProjectAllCustomer(projectName string) ([]models.Customer, error) {
	o := GetOrmer()
	//返回项目的客户列表
	//select * from customer where tag in (select label from repo_label where repoName like 'library%' group by label)
	sql := `select * from customer where tag in (select label from repo_label
	 where del=0 and repoName like ? group by label)`

	var customer []models.Customer

	if _, err := o.Raw(sql,projectName+"%").QueryRows(&customer); err != nil {
		return nil, err
	}

	log.Infof("projectName: %v", projectName)

	return customer, nil
}


func UpdateCustomerById(c models.Customer) (err error) {
	o := GetOrmer()
	res, err := o.Raw("UPDATE customer SET name = ?,tag =? WHERE id = ?",c.Name,c.Tag,c.Id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}

func DeleteCustomer(id int) (err error) {
	o := GetOrmer()
	res, err := o.Raw("delete from customer where id = ?",id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		fmt.Println("mysql row affected nums: ", num)
	}
	return err
}

