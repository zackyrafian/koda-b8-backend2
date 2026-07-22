import { useRef } from "react"
import { useEffect } from "react"
import { useState } from "react"

export default function Home() {
  const [data, setData] = useState([])
  const [open, setOpen] = useState(false)
  const inputImage = useRef(null)

  useEffect(() => { 
    const fetchData = async () => { 
      const res = await fetch("http://localhost:9999/users", {
        method: "GET",
        headers: { 
          Authorization: localStorage.getItem("token")
        }
      }) 
      const data = await res.json()
      console.log(data.results)
      setData(data.results)
    }
    fetchData()
  }, [setData])

  const handleDelete = (id) => { 
    const deleteUsers = async () => {
      const res = fetch(`http://localhost:9999/users/${id}`, { 
        method: "DELETE",
        headers: { 
          Authorization: localStorage.getItem("token")
        }
      })
      if (!res.ok) { 
        window.alert("Berhasil Menghapus")
      }
      setData((prev) => prev.filter((user) => user.id !== id));
    }
    deleteUsers()
  }

  const handleAddUsers = (e) => { 
    e.preventDefault(); 
    const form = new FormData(e.target); 
    const data = Object.fromEntries(form)
    const InsertUser = async () => { 
      const res = await fetch("http://localhost:9999/auth/sign-up", { 
        method: "POST", 
        body: JSON.stringify(data)
      })
      if (res.ok) {
        window.alert(`Berhasil Daftar, ${data.fullname}`)
        setOpen(false);
      }
    }
    InsertUser();
  }

  const handleUploadProfilePicture = (e, id) => { 
    e.preventDefault(); 
    const file = e.target.files[0]
    console.log(file)
    const form = new FormData();
    form.append('profile_picture', file)
    const fet = async() => { 
      await fetch(`http://localhost:9999/users/${id}/picture`, {
        method: "PATCH", 
        headers: { 
          Authorization: localStorage.getItem("token"),
        }, 
        body: form
      })
    }
    fet()
  }
  console.log(open);
  return (
    <div className=""> 
      {open && (
        <div className="fixed inset-0 bg-black/50 flex items-center justify-center">
          <div className="bg-white rounded-xl min-w-2xl p-4 flex flex-col gap-4">
            <div className="flex justify-between w-full">
              <div className="text-xl font-medium">Tambah Users</div>
              <button onClick={() => setOpen(prev => !prev)}>Close</button>
            </div>
            <div>
              <form onSubmit={handleAddUsers} className="flex flex-col gap-2">
                <div className="flex flex-col gap-1">
                  <label htmlFor="">FullName</label>
                  <input className="border p-1.5 rounded-lg border-black/20" type="text" name="fullname" />
                </div>
                <div className="flex flex-col gap-1">
                  <label htmlFor="">Email</label>
                  <input className="border p-1.5 rounded-lg border-black/20" type="text" name="email" />
                </div>
                <div className="flex flex-col gap-1">
                  <label htmlFor="">Password</label>
                  <input className="border p-1.5 rounded-lg border-black/20" type="password" name="password" />
                </div>
                <div className="w-full">
                  <button type="submit" className="w-full rounded-xl bg-black text-white py-2 p-1.5">Tambah</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      )}

      <div className="flex justify-between py-4 max-w-7xl mx-auto items-center">
        <h1 className="text-3xl font-medium">List Users</h1>
        <button onClick={() => setOpen(prev => !prev) } className="bg-black text-gray-200 py-2 px-4 rounded-xl w-fit text-sm">Tambah User</button>
      </div>
      
      <div className="flex flex-col max-w-7xl mx-auto items-center gap-2">
        {data.map((i) => ( 
          <div key={i.id} className="border border-black/20 w-full p-3 rounded-xl flex gap-3">
            <div onClick={() => inputImage.current.click()} className="rounded-full overflow-hidden w-12 h-12 cursor-pointer">
              <img className="w-full h-full" src={`http://localhost:9999/${i.picture}`} alt="" />
                <input ref={inputImage} className="hidden" type="file" onChange={(e) => handleUploadProfilePicture(e ,i.id)} name="profilePicture" />
            </div>
            <div className="flex gap-3 flex-col">
              <div>
                <div className="flex flex-col">
                  <span className="font-medium">Nama:  </span>
                  <span className="text-sm">{i.fullname}</span>
                </div>
                <div className="flex flex-col">
                  <span className="font-medium">Email: </span>
                  <span className="text-sm">{i.email}</span>
                </div>
              </div>
              <div className="flex gap-2 text-sm">
                <button className="px-4 py-1 text-black border-black/20 border rounded-md">Edit</button>
                <button className="px-4 bg-black text-white rounded-md" onClick={() => handleDelete(i.id)}>Hapus</button>
              </div>
            </div>
            
          </div>
        ))}
      </div>
    </div>
  )
}