import { MainPage } from "./pages/MainPage/MainPage.tsx"
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { SignIn } from "./pages/signInUpPage/SignIn.tsx";
import { SignUp } from "./pages/signInUpPage/SignUp.tsx";

export default function App() {

  return (
      <Routes>
          <Route path="/auth" index element={<SignIn />} />
          <Route path="/reg" element={<SignUp />} />
          <Route path="/" element={<MainPage />} />
      </Routes>
  );
}
