import { BrowserRouter, Routes, Route } from "react-router-dom";
import Index from "./components/pages/Index.tsx";
import AuthPage from "./components/pages/Auth.tsx";
import Dashboard from "./components/pages/Dashboard.tsx";
import AdminMain from "./components/pages/admin/AdminMain.tsx";
import CreateUser from "./components/pages/admin/CreateUser.tsx";
import DeleteUser from "./components/pages/admin/DeleteUser.tsx";
import ChangePassword from "./components/pages/admin/ChangePassword.tsx";
import ManageUser from "./components/pages/admin/ManageUser.tsx";
import Permissions from "./components/pages/admin/Permissions.tsx";
import AdminSettings from "./components/pages/admin/Settings.tsx";
import NotFoundPage from "./components/pages/errors/404.tsx";
import Navigation from "./components/base/navigation.tsx";
import { ThemeProvider } from "@/components/theme-provider";
import { Footer } from "./components/base/footer.tsx";
import UnauthorizedPage from "./components/pages/errors/401.tsx";

export default function App() {
  return (
    <BrowserRouter>
      <ThemeProvider defaultTheme="system" storageKey="vite-ui-theme">
        <Navigation />
        <Routes>
          <Route path="/" element={<Index />} />
          <Route path="/auth" element={<AuthPage />} />
          <Route path="/login" element={<AuthPage />} />
          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/admin" element={<AdminMain />} />
          <Route path="/admin/create-user" element={<CreateUser />} />
          <Route path="/admin/delete-user" element={<DeleteUser />} />
          <Route path="/admin/change-password" element={<ChangePassword />} />
          <Route path="/admin/manage-user" element={<ManageUser />} />
          <Route path="/admin/permissions" element={<Permissions />} />
          <Route path="/admin/settings" element={<AdminSettings />} />
          <Route path="/unauthorized" element={<UnauthorizedPage />} />
          <Route path="*" element={<NotFoundPage />} />
        </Routes>
      </ThemeProvider>
      <Footer />
    </BrowserRouter>
  );
}
