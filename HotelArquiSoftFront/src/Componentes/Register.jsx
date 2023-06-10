import React, {useState} from "react";

function Register(props) {

    const [email, setemail] = useState( ' ');
    const [password, setPass] = useState( ' ');
    const [name, setName] = useState( ' ');

    const handleSubmit = (

        {
            epreventDefault() {
                console.log
                {
                    email
                }

            }
        }

    )


    return (
        <Register>

            <form onSubmit={handleSubmit}>
                <label> FUll name </label>
                <input value={name} name={name} id={name} placeholder={"Full name"}/>
                <label htmlFor={"email"}>Email</label>
                <input value={email} type={"email"} placeholder={"yourEmail@gmail.com"} id={"email"} name={"email"}/>

                <label htmlFor={"Password"}>Email</label>
                <input value={password} type={"Password"} placeholder={"Hola*152"} id={"Password"} name={"Password"}/>
                <button>Log In</button>
            </form>

            <button>Ya tenes una cuenta? LOG IN</button>

        </Register>
);
}

export default Register;