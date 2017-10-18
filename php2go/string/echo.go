package string

//echo "Hello world!";

//echo "Hello world!\n";

//echo 'This ','string ','was ','made ','with multiple parameters.';

//echo 'This ','string ','was ','made ','with multiple parameters.', PHP_EOL;

func Echo(args ...string) {

	for _, v := range args {
		print(v)
	}

}
