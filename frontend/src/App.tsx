import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import ScholarshipCreate from "./components/ScholarshipCreate";
import Scholarships from "./components/Scholarships";
export default function App() {
return (
  <Router>
   <div>
   <Navbar />
   <Routes>
   <Route path="/scholarshipcreate" element={<ScholarshipCreate />} />
       <Route path="/" element={<Scholarships />} />
   </Routes>
   </div>
  </Router>
);
}
