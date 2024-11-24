import { Check, Grid, Umbrella, Circle, Activity, Globe, MapPin, Mail } from 'react-feather';

export default function Navbar() {
  return (
    <>
      <nav className="flex flex-col w-72 h-full bg-slate-800 px-6 py-5">
        <div className="flex flex-row items-center gap-2">
          <Check className="bg-blue-800 rounded-lg h-8 w-auto p-2 text-white" />
          <h1 className="text-2xl text-white">Tracked</h1>
        </div>

        <div className="mt-10 flex flex-col gap-1">
          <span className="text-sm text-slate-400 uppercase">Menu</span>
          <div className="flex flex-col gap-2">
            <div className="flex flex-row items-center gap-2">
              <Grid className="text-white w-4 h-auto" />
              <a href="#" className="text-sm text-white">Dashboard</a>
            </div>
            <div className="flex flex-row items-center gap-2">
              <Umbrella className="text-white w-4 h-auto" />
              <a href="#" className="text-sm text-white">Politicas de Privacidade</a>
            </div>
            <div className="flex flex-row items-center gap-2">
              <Circle className="text-white w-4 h-auto" />
              <a href="#" className="text-sm text-white">Cookies</a>
            </div>
            <div className="flex flex-row items-center gap-2">
              <Activity className="text-white w-4 h-auto" />
              <a href="#" className="text-sm text-white">Consentimentos</a>
            </div>
          </div>
        </div>


        <div className="mt-10 flex flex-col gap-1">
          <span className="text-sm text-slate-400 uppercase">Configuracoes </span>
          <div className="flex flex-col gap-2">
            <div className="flex flex-row items-center gap-2">
              <MapPin className="text-white w-4 h-auto" />
              <a href="#" className="text-sm text-white">Workspaces</a>
            </div>
            <div className="flex flex-row items-center gap-2">
              <Globe className="text-white w-4 h-auto" />
              <a href="#" className="text-sm text-white">Websites</a>
            </div>
          </div>
        </div>

        <div className="mt-10 flex flex-col gap-1">
          <span className="text-sm text-slate-400 uppercase">Suporte</span>
          <div className="flex flex-col gap-2">
            <div className="flex flex-row items-center gap-2">
              <Mail className="text-white w-4 h-auto" />
              <a href="#" className="text-sm text-white">Email</a>
            </div>
          </div>
        </div>

      </nav >
    </>
  )
}
