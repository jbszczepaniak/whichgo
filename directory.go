package main

type Dir struct{}

func (Dir) Validate(module string) error {
	return nil
}
func (Dir) GoVer(string) {

}
