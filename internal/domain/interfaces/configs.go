package interfaces

type TestConfig struct {
	Foo        string  `mapstructure:"foo"`
	FooInt     int32   `mapstructure:"fooInt"`
	FooFloat   float64 `mapstructure:"fooFloat"`
	FooBoolean bool    `mapstructure:"fooBoolean"`
}
