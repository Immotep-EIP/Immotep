import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import MainLayout from '@/components/MainLayout/MainLayout.tsx';
import Login from '@/views/Authentification/Login/Login.tsx';
import Register from '@/views/Authentification/Register/Register.tsx';
import OverviewContent from '@/views/Overview/Overview';
import RealPropertyContent from '@/views/RealProperty/RealProperty';
import MessagesContent from '@/views/Messages/Messages';
import Settings from '@/views/Settings/Settings';
import MyProfile from '@/views/MyProfile/MyProfile';
import NavigationEnum from '@/enums/NavigationEnum';

const App: React.FC = () => (
  <Router>
    <Routes>
      <Route path={NavigationEnum.LOGIN} element={<Login />} />
      <Route path={NavigationEnum.REGISTER} element={<Register />} />

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
