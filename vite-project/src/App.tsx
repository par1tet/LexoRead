import { MainPage } from "./pages/MainPage/MainPage.tsx"
import { Route, Routes } from "react-router-dom";
import { SignIn } from "./pages/signInUpPage/SignIn.tsx";
import { SignUp } from "./pages/signInUpPage/SignUp.tsx";
import { BooksPage } from "./pages/booksPage/BooksPage.tsx"
import { AboutUs } from "./pages/AboutUsPage/AboutUs.tsx";

export default function App() {
  return (
      <Routes>
          <Route path="/auth" index element={<SignIn />} />
          <Route path="/reg"  element={<SignUp />} />
          <Route path="/" element={<MainPage />} />
          <Route path="/books" element={<BooksPage />} />
          <Route path="/aboutUs" element={<AboutUs/>}></Route>
      </Routes>
  );
}
