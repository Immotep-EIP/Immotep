import { Route, Routes, BrowserRouter as Router } from "react-router-dom";
import "./App.css";
import HomePage from "./pages/Home/Home";
import Navbar from "./components/Navbar/Navbar";
import FeaturesPage from "./pages/Features/Features";
import OurApplicationPage from "./pages/OurApplication/OurApplication";
import PricingPage from "./pages/Pricing/Pricing";
import ContactUsPage from "./pages/ContactUs/ContactUs";
import Footer from "./components/Footer/Footer";
import NavigationEnum from "./enums/NavigationEnum";
import PrivacyPolicyPage from "./pages/Legal/PrivacyPolicy";
import LegalMentionsPage from "./pages/Legal/LegalMentions";
import DemoPage from "./pages/Demo/Demo";

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route
          path="/"
          element={
            <>
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
            </>
          }
        />
        <Route
          path={NavigationEnum.LEGAL_MENTIONS}
          element={<LegalMentionsPage />}
        />
        <Route
          path={NavigationEnum.PRIVACY_POLICY}
          element={<PrivacyPolicyPage />}
        />
        <Route path={NavigationEnum.DEMO} element={<DemoPage />} />
      </Routes>
      <Footer />
    </Router>
  );
}

export default App;
