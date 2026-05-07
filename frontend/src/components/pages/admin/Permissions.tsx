import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { Shield, ArrowLeft, Check, X } from "lucide-react";
import { useState } from "react";

interface Permission {
  id: string;
  name: string;
  description: string;
  enabled: boolean;
}

export default function Permissions() {
  const navigate = useNavigate();
  const [userId, setUserId] = useState("");
  const [login, setLogin] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState("");
  
  const [permissions, setPermissions] = useState<Permission[]>([
    { id: "admin", name: "Администратор", description: "Полный доступ ко всем функциям системы", enabled: false },
    { id: "user_management", name: "Управление пользователями", description: "Создание, удаление и редактирование пользователей", enabled: false },
    { id: "content_management", name: "Управление контентом", description: "Доступ к редактированию контента системы", enabled: false },
    { id: "reports", name: "Просмотр отчётов", description: "Доступ к аналитике и отчётам", enabled: false },
    { id: "settings", name: "Настройки системы", description: "Изменение системных настроек", enabled: false },
    { id: "api_access", name: "Доступ к API", description: "Использование API системы", enabled: false },
  ]);

  const togglePermission = (id: string) => {
    setPermissions(permissions.map(perm => 
      perm.id === id ? { ...perm, enabled: !perm.enabled } : perm
    ));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError("");
    
    // Validation
    if (!userId && !login) {
      setError("Введите ID или логин пользователя");
      setIsLoading(false);
      return;
    }
    
    const selectedPermissions = permissions.filter(p => p.enabled);
    if (selectedPermissions.length === 0) {
      setError("Выберите хотя бы одно право доступа");
      setIsLoading(false);
      return;
    }

    // Simulate API call
    setTimeout(() => {
      setIsLoading(false);
      setSuccess(true);
      setTimeout(() => {
        setSuccess(false);
        setUserId("");
        setLogin("");
        setPermissions(permissions.map(p => ({ ...p, enabled: false })));
      }, 3000);
    }, 1000);
  };

  return (
    <div className="container mx-auto px-4 py-8 max-w-4xl">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Управление правами доступа</h1>
          <p className="text-muted-foreground mt-2">
            Назначение прав доступа пользователям системы ITI
          </p>
        </div>
        <Button variant="outline" onClick={() => navigate("/admin")} className="gap-2">
          <ArrowLeft className="h-4 w-4" />
          Назад к администрированию
        </Button>
      </div>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <Shield className="h-5 w-5" />
            Назначение прав доступа
          </CardTitle>
          <CardDescription>
            Добавьте или удалите права доступа для пользователей
          </CardDescription>
        </CardHeader>
        <CardContent>
          {success ? (
            <div className="p-4 rounded-lg bg-green-50 dark:bg-green-900/20 border border-green-200 dark:border-green-800">
              <div className="flex items-center gap-2 text-green-600 dark:text-green-400">
                <div className="h-5 w-5 rounded-full bg-green-100 dark:bg-green-800 flex items-center justify-center">
                  <svg className="h-3 w-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                  </svg>
                </div>
                <span className="font-medium">Права доступа успешно обновлены!</span>
              </div>
              <p className="text-sm text-green-600 dark:text-green-400 mt-2">
                Права доступа пользователя были успешно изменены.
              </p>
            </div>
          ) : (
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="space-y-4">
                <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                  <div>
                    <label htmlFor="userId" className="block text-sm font-medium mb-2">
                      ID пользователя
                    </label>
                    <input
                      id="userId"
                      type="text"
                      value={userId}
                      onChange={(e) => setUserId(e.target.value)}
                      className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                      placeholder="Введите ID пользователя"
                    />
                    <p className="text-xs text-muted-foreground mt-1">
                      Оставьте пустым, если используете логин
                    </p>
                  </div>

                  <div>
                    <label htmlFor="login" className="block text-sm font-medium mb-2">
                      Логин пользователя
                    </label>
                    <input
                      id="login"
                      type="text"
                      value={login}
                      onChange={(e) => setLogin(e.target.value)}
                      className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                      placeholder="Введите логин пользователя"
                    />
                    <p className="text-xs text-muted-foreground mt-1">
                      Оставьте пустым, если используете ID
                    </p>
                  </div>
                </div>

                <div>
                  <h3 className="text-lg font-medium mb-4">Выберите права доступа:</h3>
                  <div className="grid grid-cols-1 md:grid-cols-2 gap-3">
                    {permissions.map((permission) => (
                      <div
                        key={permission.id}
                        className={`p-4 rounded-lg border cursor-pointer transition-colors ${
                          permission.enabled
                            ? "border-blue-500 bg-blue-50 dark:bg-blue-900/20"
                            : "border-gray-200 dark:border-gray-800 hover:border-gray-300 dark:hover:border-gray-700"
                        }`}
                        onClick={() => togglePermission(permission.id)}
                      >
                        <div className="flex items-start justify-between">
                          <div>
                            <div className="flex items-center gap-2">
                              <h4 className="font-medium">{permission.name}</h4>
                              {permission.enabled ? (
                                <Check className="h-4 w-4 text-green-500" />
                              ) : (
                                <X className="h-4 w-4 text-gray-400" />
                              )}
                            </div>
                            <p className="text-sm text-muted-foreground mt-1">
                              {permission.description}
                            </p>
                          </div>
                          <div
                            className={`h-5 w-5 rounded-full border flex items-center justify-center ${
                              permission.enabled
                                ? "bg-blue-500 border-blue-500"
                                : "border-gray-300 dark:border-gray-600"
                            }`}
                          >
                            {permission.enabled && (
                              <Check className="h-3 w-3 text-white" />
                            )}
                          </div>
                        </div>
                      </div>
                    ))}
                  </div>
                </div>

                <div className="p-4 rounded-lg bg-blue-50 dark:bg-blue-900/20 border border-blue-200 dark:border-blue-800">
                  <div className="flex items-start gap-3">
                    <Shield className="h-5 w-5 text-blue-600 dark:text-blue-400 flex-shrink-0 mt-0.5" />
                    <div>
                      <p className="text-sm font-medium text-blue-800 dark:text-blue-300">Выбранные права:</p>
                      <p className="text-sm text-blue-700 dark:text-blue-400 mt-1">
                        {permissions.filter(p => p.enabled).length === 0
                          ? "Нет выбранных прав"
                          : permissions.filter(p => p.enabled).map(p => p.name).join(", ")}
                      </p>
                    </div>
                  </div>
                </div>
              </div>

              {error && (
                <div className="p-3 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800">
                  <p className="text-sm text-red-600 dark:text-red-400">{error}</p>
                </div>
              )}

              <div className="flex flex-col sm:flex-row gap-3">
                <Button
                  type="submit"
                  disabled={isLoading || (!userId && !login)}
                  className="flex-1"
                >
                  {isLoading ? (
                    <>
                      <svg className="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" />
                        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                      </svg>
                      Сохранение...
                    </>
                  ) : (
                    "Сохранить права доступа"
                  )}
                </Button>
                <Button
                  type="button"
                  variant="outline"
                  onClick={() => navigate("/admin")}
                  className="flex-1"
                >
                  Отмена
                </Button>
              </div>
            </form>
          )}
        </CardContent>
      </Card>

      {/* API Information */}
      <Card className="mt-6">
        <CardHeader>
          <CardTitle>API Endpoint</CardTitle>
          <CardDescription>
            Backend endpoint для управления правами доступа
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="p-3 rounded-lg bg-muted">
            <code className="text-sm font-mono">POST /api/private/permissions (в разработке)</code>
            <p className="text-sm text-muted-foreground mt-1">
              Этот endpoint находится в разработке и будет доступен в будущих версиях.
            </p>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}