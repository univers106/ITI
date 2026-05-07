import { useEffect, useState } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { checkAuth, logout, getUserGreeting } from "@/lib/api";
import { useNavigate } from "react-router-dom";
import { AlertCircle, User, LogOut, Shield, Settings, Users, Key, UserPlus, UserMinus, ArrowRight } from "lucide-react";

export default function Dashboard() {
  const navigate = useNavigate();
  const [isAuthenticated, setIsAuthenticated] = useState<boolean | null>(null);
  const [userInfo, setUserInfo] = useState<{ login: string; name: string }>({ login: "", name: "" });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const verifyAuth = async () => {
      try {
        const authResult = await checkAuth();
        setIsAuthenticated(authResult);
        
        if (authResult) {
          try {
            const greeting = await getUserGreeting();
            // Parse greeting: "Hello {name}, you login is {login}"
            const nameMatch = greeting.match(/Hello\s+(.+?), you login is/);
            const loginMatch = greeting.match(/you login is\s+(.+)$/);
            
            if (nameMatch && loginMatch) {
              setUserInfo({
                name: nameMatch[1].trim(),
                login: loginMatch[1].trim()
              });
            } else {
              // Fallback if parsing fails
              setUserInfo({
                name: "Пользователь",
                login: greeting.includes("test") ? "test_user" : "user"
              });
            }
          } catch (error) {
            console.error("Failed to get user info:", error);
            // Fallback to mock data
            setUserInfo({
              login: "test_user",
              name: "Пользователь"
            });
          }
        }
      } catch (error) {
        console.error("Auth check failed:", error);
        setIsAuthenticated(false);
      } finally {
        setLoading(false);
      }
    };

    verifyAuth();
  }, []);

  const handleLogout = async () => {
    try {
      await logout();
      // После выхода перезагружаем страницу для обновления состояния
      window.location.replace("/");
    } catch (error) {
      console.error("Logout failed:", error);
    }
  };

  if (loading) {
    return (
      <div className="container mx-auto px-4 py-16 max-w-4xl">
        <div className="flex flex-col items-center justify-center min-h-[400px]">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-primary"></div>
          <p className="mt-4 text-muted-foreground">Проверка авторизации...</p>
        </div>
      </div>
    );
  }

  if (!isAuthenticated) {
    return (
      <div className="container mx-auto px-4 py-16 max-w-4xl">
        <Card>
          <CardHeader>
            <CardTitle className="text-2xl flex items-center gap-2">
              <AlertCircle className="h-6 w-6 text-destructive" />
              Доступ запрещён
            </CardTitle>
            <CardDescription>
              Для доступа к этой странице необходимо войти в систему
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-4">
            <p>Вы не авторизованы для просмотра этой страницы.</p>
            <div className="flex gap-3">
              <Button onClick={() => navigate("/auth")}>
                Войти в систему
              </Button>
              <Button variant="outline" onClick={() => navigate("/")}>
                На главную
              </Button>
            </div>
          </CardContent>
        </Card>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-8 max-w-4xl">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Личный кабинет</h1>
          <p className="text-muted-foreground mt-2">
            Вся информация о вашем аккаунте
          </p>
        </div>
        <div className="flex flex-col sm:flex-row gap-3">
          <Button variant="outline" onClick={() => navigate("/admin")} className="gap-2">
            <Settings className="h-4 w-4" />
            Администрирование
            <ArrowRight className="h-4 w-4" />
          </Button>
          <Button variant="outline" onClick={handleLogout} className="gap-2">
            <LogOut className="h-4 w-4" />
            Выйти
          </Button>
        </div>
      </div>

      {/* User Information */}
      <Card className="mb-8">
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <User className="h-5 w-5" />
            Информация о пользователе
          </CardTitle>
          <CardDescription>
            Все данные, которые вы можете знать о своём аккаунте
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div className="space-y-4">
              <div>
                <p className="text-sm font-medium text-muted-foreground mb-1">Логин</p>
                <p className="text-lg font-semibold">{userInfo.login || "Не указан"}</p>
              </div>
              <div>
                <p className="text-sm font-medium text-muted-foreground mb-1">Имя</p>
                <p className="text-lg font-semibold">{userInfo.name || "Не указано"}</p>
              </div>
              <div>
                <p className="text-sm font-medium text-muted-foreground mb-1">Статус аккаунта</p>
                <div className="flex items-center gap-2">
                  <div className="h-2 w-2 rounded-full bg-green-500"></div>
                  <span className="text-green-600 font-medium">Активен</span>
                </div>
              </div>
            </div>
            
            <div className="space-y-4">
              <div>
                <p className="text-sm font-medium text-muted-foreground mb-1">Последний вход</p>
                <p className="text-lg font-semibold">Сегодня</p>
              </div>
              <div>
                <p className="text-sm font-medium text-muted-foreground mb-1">Тип сессии</p>
                <p className="text-lg font-semibold">Веб-браузер</p>
              </div>
              <div>
                <p className="text-sm font-medium text-muted-foreground mb-1">Доступ к системе</p>
                <div className="flex items-center gap-2">
                  <Shield className="h-4 w-4 text-green-500" />
                  <span className="text-green-600 font-medium">Разрешён</span>
                </div>
              </div>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Administration Section */}
      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Settings className="h-5 w-5" />
            Администрирование
          </CardTitle>
          <CardDescription>
            Управление пользователями системы (доступно администраторам)
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            <Button 
              variant="outline" 
              className="h-auto py-4 flex flex-col gap-2 items-center justify-center"
              onClick={() => alert("Переход к созданию пользователя\nAPI endpoint: POST /api/private/user-manipulation/create")}
            >
              <UserPlus className="h-6 w-6" />
              <div>
                <div className="font-medium">Создать пользователя</div>
                <div className="text-xs text-muted-foreground font-normal mt-1">
                  POST /user-manipulation/create
                </div>
              </div>
            </Button>
            
            <Button 
              variant="outline" 
              className="h-auto py-4 flex flex-col gap-2 items-center justify-center"
              onClick={() => alert("Переход к удалению пользователя\nAPI endpoint: POST /api/private/user-manipulation/delete")}
            >
              <UserMinus className="h-6 w-6" />
              <div>
                <div className="font-medium">Удалить пользователя</div>
                <div className="text-xs text-muted-foreground font-normal mt-1">
                  POST /user-manipulation/delete
                </div>
              </div>
            </Button>
            
            <Button 
              variant="outline" 
              className="h-auto py-4 flex flex-col gap-2 items-center justify-center"
              onClick={() => alert("Переход к изменению пароля\nAPI endpoint: POST /api/private/user-manipulation/change-password")}
            >
              <Key className="h-6 w-6" />
              <div>
                <div className="font-medium">Изменить пароль</div>
                <div className="text-xs text-muted-foreground font-normal mt-1">
                  POST /user-manipulation/change-password
                </div>
              </div>
            </Button>
            
            <Button 
              variant="outline" 
              className="h-auto py-4 flex flex-col gap-2 items-center justify-center"
              onClick={() => alert("Переход к изменению логина\nAPI endpoint: POST /api/private/user-manipulation/change-login")}
            >
              <User className="h-6 w-6" />
              <div>
                <div className="font-medium">Изменить логин</div>
                <div className="text-xs text-muted-foreground font-normal mt-1">
                  POST /user-manipulation/change-login
                </div>
              </div>
            </Button>
            
            <Button 
              variant="outline" 
              className="h-auto py-4 flex flex-col gap-2 items-center justify-center"
              onClick={() => alert("Переход к изменению имени\nAPI endpoint: POST /api/private/user-manipulation/change-name")}
            >
              <Users className="h-6 w-6" />
              <div>
                <div className="font-medium">Изменить имя</div>
                <div className="text-xs text-muted-foreground font-normal mt-1">
                  POST /user-manipulation/change-name
                </div>
              </div>
            </Button>
            
            <Button 
              variant="outline" 
              className="h-auto py-4 flex flex-col gap-2 items-center justify-center"
              onClick={() => alert("Переход к управлению правами\nAPI endpoints доступны в бэкенде")}
            >
              <Shield className="h-6 w-6" />
              <div>
                <div className="font-medium">Управление правами</div>
                <div className="text-xs text-muted-foreground font-normal mt-1">
                  Database permissions API
                </div>
              </div>
            </Button>
          </div>
          
          <div className="mt-6 p-4 bg-muted rounded-lg">
            <h4 className="font-medium mb-2">Доступные API endpoints для администрирования:</h4>
            <ul className="text-sm space-y-1 text-muted-foreground">
              <li>• POST /api/private/user-manipulation/create - Создание пользователя</li>
              <li>• POST /api/private/user-manipulation/delete - Удаление пользователя</li>
              <li>• POST /api/private/user-manipulation/change-password - Изменение пароля</li>
              <li>• POST /api/private/user-manipulation/change-login - Изменение логина</li>
              <li>• POST /api/private/user-manipulation/change-name - Изменение имени</li>
              <li>• Database.UserAddPermissions - Добавление прав пользователю</li>
              <li>• Database.UserRemovePermissions - Удаление прав пользователя</li>
            </ul>
          </div>
        </CardContent>
      </Card>

      {/* Footer */}
      <div className="mt-8 text-center text-sm text-muted-foreground">
        <p>Система ITI • {new Date().getFullYear()} ©</p>
        <p className="mt-1">Для управления пользователями используйте кнопки выше.</p>
      </div>
    </div>
  );
}