typedef int(*intFunc) ();

int bridge_int_func(intFunc f)
{
	return f();
}

int fortytwo()
{
	return 42;
}