import { MainPage } from "./pages/MainPage/MainPage.tsx"
import { Routes, Route } from "react-router-dom"

function App() {
  return (
    <Routes>
        <Route path="/" element={<MainPage />} />
    </Routes>
  )
}

export default App
