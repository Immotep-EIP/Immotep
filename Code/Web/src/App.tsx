import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import NavigationEnum from '@/enums/NavigationEnum';
import MainLayout from '@/components/MainLayout/MainLayout.tsx';

// ! AUTHENTIFICATION
import Login from '@/views/Authentification/Login/Login.tsx';
import Register from '@/views/Authentification/Register/Register.tsx';
import ForgotPassword from '@/views/Authentification/ForgotPassword/ForgotPassword.tsx';

// ! MAIN LAYOUT - SIDEBAR
import OverviewContent from '@/views/Overview/Overview';
import RealPropertyContent from '@/views/RealProperty/RealProperty';
import MessagesContent from '@/views/Messages/Messages';

// ! MAIN LAYOUT - HEADER
import Settings from '@/views/Settings/Settings';
import MyProfile from '@/views/MyProfile/MyProfile';


const App: React.FC = () => (
  <Router>
    <Routes>
      <Route path={NavigationEnum.LOGIN} element={<Login />} />
      <Route path={NavigationEnum.REGISTER} element={<Register />} />
      <Route path={NavigationEnum.FORGOT_PASSWORD} element={<ForgotPassword />} />

      <Route element={<MainLayout />}>
        <Route path={NavigationEnum.OVERVIEW} element={<OverviewContent />} />
        <Route path={NavigationEnum.REAL_PROPERTY} element={<RealPropertyContent />} />
        <Route path={NavigationEnum.MESSAGES} element={<MessagesContent />} />
        <Route path={NavigationEnum.SETTINGS} element={<Settings />} />
        <Route path={NavigationEnum.MY_PROFILE} element={<MyProfile />} />
      </Route>
    </Routes>
  </Router>
);

export default App;
