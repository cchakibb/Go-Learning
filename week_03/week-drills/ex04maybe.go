type Notifier struct {
    Name string
    Send func(message string) error  // champ fonction
}

func (n Notifier) Label() string {  // méthode
    return "[" + n.Name + "]"
}

type Sender interface {
    Send(message string) error  // méthode d'interface
}

func notify(s Sender, msg string) {
    s.Send(msg)  // appel A
}

func main() {
    n := Notifier{
        Name: "email",
        Send: func(msg string) error {
            fmt.Println("sending:", msg)
            return nil
        },
    }

    n.Send("hello")   // appel B
    n.Label()         // appel C
    notify(n, "hi")   // appel D
}