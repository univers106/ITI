import { Link, useNavigate, useLocation } from "react-router-dom";
import { useState, useEffect } from "react";
import { checkAuth, logout } from "@/lib/api";
import { Button } from "@/components/ui/button";
import { LogOut, LogIn } from "lucide-react";

const NAV_ITEMS = [
  { label: "Главная", href: "/" },
  { label: "", href: "" },
  { label: "О нас", href: "/about" },
  { label: "Контакты", href: "/contacts" },
  { label: "Политика конфиденциальности", href: "/privacy" },
] as const;

export function Footer() {
  const navigate = useNavigate();
  const location = useLocation();
  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const verifyAuth = async () => {
      try {
        const authResult = await checkAuth();
        setIsAuthenticated(authResult);
      } catch (error) {
        console.error("Auth check failed:", error);
        setIsAuthenticated(false);
      } finally {
        setLoading(false);
      }
    };

    verifyAuth();
    
    // Перепроверяем периодически
    const interval = setInterval(verifyAuth, 30000); // Каждые 30 секунд
    return () => clearInterval(interval);
  }, [location.pathname]); // Перепроверяем при смене пути

  const handleAuthAction = async () => {
    if (isAuthenticated) {
      try {
        await logout();
        // После выхода перезагружаем страницу для обновления состояния
        window.location.replace("/");
      } catch (error) {
        console.error("Logout failed:", error);
      }
    } else {
      navigate("/auth");
    }
  };

  const getAuthButtonText = () => {
    if (loading) return "Загрузка...";
    return isAuthenticated ? "Выйти из аккаунта" : "Вход в аккаунт";
  };

  const getAuthIcon = () => {
    if (loading) return null;
    return isAuthenticated ? <LogOut className="h-4 w-4" /> : <LogIn className="h-4 w-4" />;
  };

  return (
    <footer className="bg-gray-100 dark:bg-gray-900 py-8">
      <div className="container mx-auto px-4 md:px-8 lg:px-16">
        <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-6">
          <div className="flex flex-col items-start gap-1">
            {NAV_ITEMS.map((item, index) => (
              item.label ? (
                <Link
                  key={index}
                  to={item.href}
                  className="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200 transition-colors"
                >
                  {item.label}
                </Link>
              ) : (
                <div key={index} className="h-4"></div>
              )
            ))}
          </div>
          
          <div className="flex flex-col items-start md:items-end gap-4">
            <Button
              variant={isAuthenticated ? "destructive" : "default"}
              size="sm"
              onClick={handleAuthAction}
              disabled={loading}
              className="gap-2"
            >
              {getAuthIcon()}
              {getAuthButtonText()}
            </Button>
            
            {isAuthenticated && (
              <Link
                to="/dashboard"
                className="text-sm text-gray-600 dark:text-gray-400 hover:text-gray-900 dark:hover:text-gray-200 transition-colors"
              >
                Личный кабинет
              </Link>
            )}
            
            <p className="text-xs text-gray-500 dark:text-gray-500 mt-2">
              © {new Date().getFullYear()} ИТИ 2 версия!!. Все права защищены.
            </p>
          </div>
        </div>
      </div>
    </footer>
  );
}
