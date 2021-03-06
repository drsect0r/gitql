package main

type CmdQuery struct {
	cmdQueryBase

	Format string `short:"f" long:"format" default:"pretty" description:"Ouptut format. Formats supported: pretty, csv, json."`
	Args   struct {
		SQL string `positional-arg-name:"sql" required:"true" description:"SQL query to execute"`
	} `positional-args:"yes"`
}

func (c *CmdQuery) Execute(args []string) error {
	if err := c.validate(); err != nil {
		return err
	}

	if err := c.buildDatabase(); err != nil {
		return err
	}

	schema, rowIter, err := c.executeQuery(c.Args.SQL)
	if err != nil {
		return err
	}

	return c.printQuery(schema, rowIter, c.Format)
}
