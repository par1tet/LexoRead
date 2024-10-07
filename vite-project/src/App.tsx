import { BrowserRouter, Route, Routes } from "react-router-dom";
import SignIn from "./pages/signInUpPage/SignIn";
import SignUp from "./pages/signInUpPage/SignUp";
export default function App() {

  return (
    <BrowserRouter>
      <Routes>
        
          <Route path="/auth" index element={<SignIn />} />
          <Route path="/reg" element={<SignUp />} />
        
      </Routes>
    </BrowserRouter>
  );
}
