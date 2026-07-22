"use client";

export default function RegisterPage() { 
  const handleSubmit = (e) => { 
    e.preventDefault();
    const form = new FormData(e.target); 
    const data = Object.fromEntries(form);
    const handleLogin = async () => {
      try { 
        const res = await fetch("http://localhost:9999/auth/sign-up",{ 
          method: "POST", 
          body: JSON.stringify(data),
        })
        if (res.ok) {
          window.alert(`Berhasil Daftar, ${data.fullname}`)
          window.location.href = '/login'
        }
      } catch (err) { 
        window.alert(`Gagal Register`)
        console.log(err)
      }
    }
    handleLogin();
  }
  return (
    <div className="flex">
      <div className="w-1/2 bg-gray-500 flex justify-center flex-col gap-4 px-4">
        <div className="flex flex-col gap-4">
          <div className="text-7xl font-medium">Mari Bergabung Dengan Website Phising <br /> Kami</div>
        </div>
        <div>
          <div className="text-2xl">Tujuan Utama Kami adalah mengambil akun FreeFire anda.</div>
          <div className="text-xl">Kami sudah mencuri 100 Rb akun freefire</div>
        </div>
      </div>
      <div className="flex-1 px-30 my-auto min-h-screen flex flex-col justify-center gap-4">
        <h1 className="text-4xl font-bold">Register</h1>
        <div className="flex gap-2">
          <div className="w-1/2 p-2 h-10 bg-black rounded-md text-white flex justify-center items-center">Google</div>
          <div className="flex-1 p-2 h-10 bg-black rounded-md text-white flex justify-center items-center">FreeFire</div>
        </div>
        <form className="flex flex-col gap-2"  onSubmit={handleSubmit}>
          <div className="flex flex-col gap-2">
            <label htmlFor="" className="text-sm">FullName: </label>
            <input className="border border-black/20 rounded-lg p-2" type="text" name="fullname" />
          </div>
          <div className="flex flex-col gap-2">
            <label htmlFor="" className="text-sm">Email: </label>
            <input className="border border-black/20 rounded-lg p-2" type="email" name="email" />
          </div>
          <div className="flex flex-col gap-2">
            <label htmlFor="" className="text-sm">Password</label>
            <input className="border border-black/20 rounded-lg p-2" type="password" name="password" />
          </div>
          <button className="bg-black text-white px-4 w-full rounded-lg my-4 p-2" type="submit">Register</button>
        </form>
      </div>
    </div>
    
  )
}