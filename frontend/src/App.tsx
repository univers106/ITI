import { BrowserRouter, Routes, Route } from "react-router-dom";
import Index from "./components/pages/Index.tsx";
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
          <Route path="*" element={<NotFoundPage />} />
          <Route path="/unauthorized" element={<UnauthorizedPage />} />
        </Routes>
      </ThemeProvider>
      <Footer />
    </BrowserRouter>
  );
}
