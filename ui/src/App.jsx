import { Umbrella, Circle, Activity, User } from 'react-feather';
import Navbar from './components/navbar';
import TopNavbar from './components/topnavbar';

function App() {
  return (
    <>
      <div className="flex h-screen overflow-hidden bg-slate-100">
        <Navbar />

        <div className="w-full">
          <TopNavbar />

          <div className="p-4">
            <div className="flex flex-wrap gap-4 justify-center">
              <div className="bg-white drop-shadow rounded-sm p-4 w-[calc(25%-1rem)]">
                <Activity className="rounded-full bg-slate-200 p-2 h-10 w-auto mb-2 text-blue-800" />
                <span className="text-xl text-slate-800 font-bold">826k</span>
                <h2 className="text-sm text-slate-400" >Consentimentos</h2>
              </div>
              <div className="bg-white drop-shadow rounded-sm p-4 w-[calc(25%-1rem)]">
                <Circle className="rounded-full bg-slate-200 p-2 h-10 w-auto mb-2 text-blue-800" />
                <span className="text-xl text-slate-800 font-bold">23</span>
                <h2 className="text-sm text-slate-400" >Cookies</h2>
              </div>
              <div className="bg-white drop-shadow rounded-sm p-4 w-[calc(25%-1rem)]">
                <Umbrella className="rounded-full bg-slate-200 p-2 h-10 w-auto mb-2 text-blue-800" />
                <span className="text-xl text-slate-800 font-bold">2</span>
                <h2 className="text-sm text-slate-400" >Politicas de privacidade</h2>
              </div>
              <div className="bg-white drop-shadow rounded-sm p-4 w-[calc(25%-1rem)]">
                <User className="rounded-full bg-slate-200 p-2 h-10 w-auto mb-2 text-blue-800" />
                <span className="text-xl text-slate-800 font-bold">129k</span>
                <h2 className="text-sm text-slate-400" >Usuarios unicos</h2>
              </div>
            </div>
          </div>
        </div >
      </div >
    </>
  )
}

export default App
