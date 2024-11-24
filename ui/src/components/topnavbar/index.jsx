export default function Navbar() {
  return (
    <>
      <nav className="bg-white flex flex-row gap-2 drop-shadow-sm py-5 px-4">
        <select className="border border-slate-200 rounded-md p-2">
          <option>Octomoney</option>
        </select>
        <select className="border border-slate-200 rounded-md p-2">
          <option>www.octomoney.com.br</option>
        </select>
      </nav>
    </>
  )
}
