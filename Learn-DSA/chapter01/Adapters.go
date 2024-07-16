package chapter01

import "fmt"

type Mobile interface {
	ChargeAppleMobile()
}

type apple struct{}

func (a *apple) ChargeAppleMobile() {
	fmt.Println("Charging Apple Mobile")
}

//Adaptee
type android struct{}

func (a *android) ChargeAndroidMobile() {
	fmt.Println("Charging Android Mobile")
}

// Adapter
type androidAdapter struct {
	android *android
}

func (ad *androidAdapter) ChargeAppleMobile() {
	ad.android.ChargeAndroidMobile()
}

type Client struct{}

func (c *Client) ChargeMobile(mob Mobile) {
	mob.ChargeAppleMobile()
}

func AdapterDSA() {
	fmt.Println("Adapter DSA")

	apple := &apple{}
	Client := &Client{}
	Client.ChargeMobile(apple)

	//Extended requiremet : charge mobile

	android := &android{}

	androidAdapter := &androidAdapter{
		android: android,
	}
	Client.ChargeMobile(androidAdapter)
}
