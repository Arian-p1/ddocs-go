package main

import (
	"fmt"
	"os"

	"github.com/Arian-p1/ddocs-go/internal"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ddocs"
	doc := `ddocs
  search <SEARCH>  search the topic you saved
  add <ADD>        add your new topic or replace with old one
  cat <CAT>        cat your topic
  edit <EDIT>      edit topic (its on TODO)
  delete <DELETE>  delete the topic
  help             Print help`
	app.Action = func(c *cli.Context) error {
		arg := c.Args().Get(0)
		switch arg {
		case "search":
			search(c.Args().Get(1))
		case "add":
			add(c.Args().Get(1))
		case "cat":
			r, _ := cat(c.Args().Get(1))
			fmt.Println(r)
		case "edit":
			edit(c.Args().Get(1))
		case "delete":
			delete_data(c.Args().Get(1))
		default:
			fmt.Println(doc)
		}
		return nil
	}
	app.Run(os.Args)
}

func add(key string) error {
	datamap, err := internal.ReadFile()
	if err != nil {
		return err
	}
	new_value, err := internal.Editor("")
	compress_value, err := internal.Compress(new_value)
	if err != nil {
		return err
	}
	datamap[key] = compress_value
	err = internal.WriteToFile(datamap)
	if err != nil {
		return err
	}
	return nil
}

func edit(key string) error {
	datamap, err := internal.ReadFile()
	if err != nil {
		return err
	}
	old, err := internal.Decompress(datamap[key])
	new_data, err := internal.Editor(old)
	c_new_data, err := internal.Compress(new_data)
	datamap[key] = c_new_data
	err = internal.WriteToFile(datamap)
	if err != nil {
		return err
	}
	return nil
}

func cat(key string) (string, error) {
	datamap, err := internal.ReadFile()
	if err != nil {
		return "", err
	}
	data, err := internal.Decompress(datamap[key])
	if err != nil {
		return "", err
	}
	return data, nil
}

func delete_data(key string) error {
	datamap, err := internal.ReadFile()
	if err != nil {
		return err
	}
	delete(datamap, key)
	err = internal.WriteToFile(datamap)
	if err != nil {
		return err
	}
	return nil
}

func search(key string) string {
	data_map, _ := internal.ReadFile()
	return data_map[key]
}
