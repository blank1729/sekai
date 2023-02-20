import React from "react";

function Login() {
  const handleSubmit = (e : React.FormEvent) => {
    e.preventDefault()
  }
  return <form className="grid w-[300px] m-auto" onSubmit={handleSubmit}>
    <label htmlFor="email">Email</label>
    <input id="input" type="email" />
    <label htmlFor="Passowrd">Password</label>
    <input id="input" type="password" />
    <button className="bg-blue-400 px-2 py-1">Submit</button>
  </form>;
}

export default Login;
