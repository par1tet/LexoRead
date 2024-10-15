import { MainPage } from "./pages/MainPage/MainPage.tsx"
import { Route, Routes } from "react-router-dom";
import { SignIn } from "./pages/signInUpPage/SignIn.tsx";
import { SignUp } from "./pages/signInUpPage/SignUp.tsx";
import { BookPage } from './pages/BookPage/BookPage.tsx'
import { AboutUs } from "./pages/AboutUsPage/AboutUs.tsx";

export default function App() {
  return (
      <Routes>
          <Route path="/auth" index element={<SignIn />} />
          <Route path="/reg"  element={<SignUp />} />
          <Route path="/" element={<MainPage />} />
          <Route path="/book" element={<BookPage />} />
          <Route path="/aboutUs" element={<AboutUs/>}></Route>
      </Routes>
  );
}
