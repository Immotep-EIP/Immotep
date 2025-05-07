import { BrowserRouter as Router } from "react-router-dom";
import HomePage from "./pages/Home/Home";
import "./App.css";
import Navbar from "./components/Navbar/Navbar";
import FeaturesPage from "./pages/Features/Features";
import OurApplicationPage from "./pages/OurApplication/OurApplication";
import PricingPage from "./pages/Pricing/Pricing";
import ContactUsPage from "./pages/ContactUs/ContactUs";
import Footer from "./components/Footer/Footer";

function App() {
  return (
    <Router>
      <Navbar />
      <div id="home">
        <HomePage />
      </div>
      <div id="features">
        <FeaturesPage />
      </div>
      <div id="application">
        <OurApplicationPage />
      </div>
      <div id="pricing">
        <PricingPage />
      </div>
      <div id="contact-us">
        <ContactUsPage />
      </div>
      <Footer />
      {/* <button onClick={() => changeLanguage("en")}>EN</button>
      <button onClick={() => changeLanguage("fr")}>FR</button>
      <button onClick={() => changeLanguage("de")}>DE</button> */}
    </Router>
  );
}

export default App;
