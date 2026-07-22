"use client";

import { useNavigate } from "react-router";

export default function LoginPage() { 
  const navigate = useNavigate()
  const handleSubmit = (e) => { 
    e.preventDefault();
    const form = new FormData(e.target); 
    const data = Object.fromEntries(form);
    const handleLogin = async () => {
      const res = await fetch("http://localhost:9999/auth/sign-in",{ 
        method: "POST", 
        body: JSON.stringify(data),
      })
      const dataResult = await res.text(); 
      const token = JSON.parse(dataResult)
      console.log(dataResult);
      localStorage.setItem("token", "Bearer " + token)
      if (res.ok) { 
        alert("Login berhasil")
        navigate("/")
      }
    }
    handleLogin();
  }
  return (
    <div>
      <form onSubmit={handleSubmit}>
        <div>
          <label htmlFor="">Email</label>
          <input type="email" name="email"/>
        </div>
        <div>
          <label htmlFor="">Password</label>
          <input type="password" name="password" />
        </div>

        <button type="submit">Submit</button>
      </form>
    </div>
  )
}