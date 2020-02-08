package main


type GoMod struct{}

func (GoMod) Validate(module string) error {
	return nil
}
func (GoMod) GoVer(string) {

}
