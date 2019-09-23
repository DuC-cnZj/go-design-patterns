package option

type Options struct {
	Name string
	Id int
	Age int64
}

type Optional func(*Options)

func SetName(name string) Optional  {
	return func(options *Options) {
		options.Name = name
	}
}

func SetId(id int) Optional  {
	return func(options *Options) {
		options.Id = id
	}
}

func SetAge(age int64) Optional  {
	return func(options *Options) {
		options.Age = age
	}
}

func New(optionals ...Optional) *Options {
	O := &Options{
		Name: "",
		Id:   0,
		Age:  0,
	}

	for _, setter := range optionals {
		setter(O)
	}

	return O
}