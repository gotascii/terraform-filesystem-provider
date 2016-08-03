package main

import (
	"errors"
	"io/ioutil"

	"github.com/hashicorp/terraform/helper/schema"
)

// ResourceLocalFile exposes a Local File Resource
func ResourceLocalFile() *schema.Resource {
	return &schema.Resource{
		Create: create,
		Read:   read,
		Update: update,
		Delete: del,

		Schema: map[string]*schema.Schema{
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"contents": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func create(d *schema.ResourceData, m interface{}) (e error) {
	ioutil.WriteFile(d.Get("path").(string), []byte(d.Get("contents").(string)), 0666)
	d.SetId(d.Get("path").(string))
	return nil
}

func read(d *schema.ResourceData, m interface{}) (e error) {
	data, _ := ioutil.ReadFile(d.Get("path").(string))
	s := string(data[:])

	d.Set("path", d.Get("path").(string))
	d.Set("contents", s)

	return nil
}

func update(d *schema.ResourceData, m interface{}) (e error) {
	d.Partial(true)
	return errors.New("FAIL")
}

func del(d *schema.ResourceData, m interface{}) (e error) {
	return nil
}
