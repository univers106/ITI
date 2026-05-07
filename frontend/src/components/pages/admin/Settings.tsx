import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { Settings as SettingsIcon, ArrowLeft, Save } from "lucide-react";
import { useState } from "react";

export default function AdminSettings() {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState("");
  
  const [systemSettings, setSystemSettings] = useState({
    sessionTimeout: 30,
    maxLoginAttempts: 5,
    passwordMinLength: 6,
    requireTwoFactor: false,
    enableAuditLog: true,
    maintenanceMode: false,
    apiRateLimit: 100,
    emailNotifications: true,
  });

  const handleSettingChange = (key: keyof typeof systemSettings, value: any) => {
    setSystemSettings(prev => ({ ...prev, [key]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError("");
    
    // Simulate API call
    setTimeout(() => {
      setIsLoading(false);
      setSuccess(true);
      setTimeout(() => {
        setSuccess(false);
      }, 3000);
    }, 1000);
  };

  return (
    <div className="container mx-auto px-4 py-8 max-w-4xl">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Настройки системы</h1>
          <p className="text-muted-foreground mt-2">
            Общие настройки системы ITI
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
            <SettingsIcon className="h-5 w-5" />
            Конфигурация системы
          </CardTitle>
          <CardDescription>
            Настройки конфигурации и параметров системы
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
                <span className="font-medium">Настройки успешно сохранены!</span>
              </div>
              <p className="text-sm text-green-600 dark:text-green-400 mt-2">
                Настройки системы были успешно обновлены.
              </p>
            </div>
          ) : (
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                {/* Security Settings */}
                <div className="space-y-4">
                  <h3 className="text-lg font-medium">Безопасность</h3>
                  
                  <div>
                    <label htmlFor="sessionTimeout" className="block text-sm font-medium mb-2">
                      Таймаут сессии (минут)
                    </label>
                    <input
                      id="sessionTimeout"
                      type="number"
                      min="1"
                      max="1440"
                      value={systemSettings.sessionTimeout}
                      onChange={(e) => handleSettingChange("sessionTimeout", parseInt(e.target.value))}
                      className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                    <p className="text-xs text-muted-foreground mt-1">
                      Время бездействия до автоматического выхода
                    </p>
                  </div>

                  <div>
                    <label htmlFor="maxLoginAttempts" className="block text-sm font-medium mb-2">
                      Максимум попыток входа
                    </label>
                    <input
                      id="maxLoginAttempts"
                      type="number"
                      min="1"
                      max="20"
                      value={systemSettings.maxLoginAttempts}
                      onChange={(e) => handleSettingChange("maxLoginAttempts", parseInt(e.target.value))}
                      className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                    <p className="text-xs text-muted-foreground mt-1">
                      Количество неудачных попыток до блокировки
                    </p>
                  </div>

                  <div>
                    <label htmlFor="passwordMinLength" className="block text-sm font-medium mb-2">
                      Минимальная длина пароля
                    </label>
                    <input
                      id="passwordMinLength"
                      type="number"
                      min="4"
                      max="32"
                      value={systemSettings.passwordMinLength}
                      onChange={(e) => handleSettingChange("passwordMinLength", parseInt(e.target.value))}
                      className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                  </div>

                  <div className="flex items-center justify-between">
                    <div>
                      <label htmlFor="requireTwoFactor" className="block text-sm font-medium">
                        Требовать двухфакторную аутентификацию
                      </label>
                      <p className="text-xs text-muted-foreground">
                        Для всех административных аккаунтов
                      </p>
                    </div>
                    <div className="relative">
                      <input
                        id="requireTwoFactor"
                        type="checkbox"
                        checked={systemSettings.requireTwoFactor}
                        onChange={(e) => handleSettingChange("requireTwoFactor", e.target.checked)}
                        className="sr-only"
                      />
                      <div
                        className={`h-6 w-11 rounded-full transition-colors cursor-pointer ${
                          systemSettings.requireTwoFactor ? "bg-blue-500" : "bg-gray-300 dark:bg-gray-700"
                        }`}
                        onClick={() => handleSettingChange("requireTwoFactor", !systemSettings.requireTwoFactor)}
                      >
                        <div
                          className={`h-5 w-5 rounded-full bg-white transform transition-transform mt-0.5 ${
                            systemSettings.requireTwoFactor ? "translate-x-5" : "translate-x-0.5"
                          }`}
                        />
                      </div>
                    </div>
                  </div>
                </div>

                {/* System Settings */}
                <div className="space-y-4">
                  <h3 className="text-lg font-medium">Система</h3>
                  
                  <div>
                    <label htmlFor="apiRateLimit" className="block text-sm font-medium mb-2">
                      Лимит запросов API (в минуту)
                    </label>
                    <input
                      id="apiRateLimit"
                      type="number"
                      min="10"
                      max="1000"
                      value={systemSettings.apiRateLimit}
                      onChange={(e) => handleSettingChange("apiRateLimit", parseInt(e.target.value))}
                      className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                    />
                    <p className="text-xs text-muted-foreground mt-1">
                      Максимальное количество запросов к API
                    </p>
                  </div>

                  <div className="flex items-center justify-between">
                    <div>
                      <label htmlFor="enableAuditLog" className="block text-sm font-medium">
                        Включить аудит-лог
                      </label>
                      <p className="text-xs text-muted-foreground">
                        Запись всех административных действий
                      </p>
                    </div>
                    <div className="relative">
                      <input
                        id="enableAuditLog"
                        type="checkbox"
                        checked={systemSettings.enableAuditLog}
                        onChange={(e) => handleSettingChange("enableAuditLog", e.target.checked)}
                        className="sr-only"
                      />
                      <div
                        className={`h-6 w-11 rounded-full transition-colors cursor-pointer ${
                          systemSettings.enableAuditLog ? "bg-blue-500" : "bg-gray-300 dark:bg-gray-700"
                        }`}
                        onClick={() => handleSettingChange("enableAuditLog", !systemSettings.enableAuditLog)}
                      >
                        <div
                          className={`h-5 w-5 rounded-full bg-white transform transition-transform mt-0.5 ${
                            systemSettings.enableAuditLog ? "translate-x-5" : "translate-x-0.5"
                          }`}
                        />
                      </div>
                    </div>
                  </div>

                  <div className="flex items-center justify-between">
                    <div>
                      <label htmlFor="emailNotifications" className="block text-sm font-medium">
                        Email уведомления
                      </label>
                      <p className="text-xs text-muted-foreground">
                        Отправлять уведомления на email
                      </p>
                    </div>
                    <div className="relative">
                      <input
                        id="emailNotifications"
                        type="checkbox"
                        checked={systemSettings.emailNotifications}
                        onChange={(e) => handleSettingChange("emailNotifications", e.target.checked)}
                        className="sr-only"
                      />
                      <div
                        className={`h-6 w-11 rounded-full transition-colors cursor-pointer ${
                          systemSettings.emailNotifications ? "bg-blue-500" : "bg-gray-300 dark:bg-gray-700"
                        }`}
                        onClick={() => handleSettingChange("emailNotifications", !systemSettings.emailNotifications)}
                      >
                        <div
                          className={`h-5 w-5 rounded-full bg-white transform transition-transform mt-0.5 ${
                            systemSettings.emailNotifications ? "translate-x-5" : "translate-x-0.5"
                          }`}
                        />
                      </div>
                    </div>
                  </div>

                  <div className="flex items-center justify-between">
                    <div>
                      <label htmlFor="maintenanceMode" className="block text-sm font-medium">
                        Режим обслуживания
                      </label>
                      <p className="text-xs text-muted-foreground">
                        Временно отключить систему для всех пользователей
                      </p>
                    </div>
                    <div className="relative">
                      <input
                        id="maintenanceMode"
                        type="checkbox"
                        checked={systemSettings.maintenanceMode}
                        onChange={(e) => handleSettingChange("maintenanceMode", e.target.checked)}
                        className="sr-only"
                      />
                      <div
                        className={`h-6 w-11 rounded-full transition-colors cursor-pointer ${
                          systemSettings.maintenanceMode ? "bg-amber-500" : "bg-gray-300 dark:bg-gray-700"
                        }`}
                        onClick={() => handleSettingChange("maintenanceMode", !systemSettings.maintenanceMode)}
                      >
                        <div
                          className={`h-5 w-5 rounded-full bg-white transform transition-transform mt-0.5 ${
                            systemSettings.maintenanceMode ? "translate-x-5" : "translate-x-0.5"
                          }`}
                        />
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              {error && (
                <div className="p-3 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800">
                  <p className="text-sm text-red-600 dark:text-red-400">{error}</p>
                </div>
              )}

              <div className="flex flex-col sm:flex-row gap-3 pt-4 border-t">
                <Button
                  type="submit"
                  disabled={isLoading}
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
                    <>
                      <Save className="h-4 w-4 mr-2" />
                      Сохранить настройки
                    </>
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

      {/* Danger Zone */}
      <Card className="mt-6 border-red-200 dark:border-red-800">
        <CardHeader>
          <CardTitle className="text-red-600 dark:text-red-400">Опасная зона</CardTitle>
          <CardDescription>
            Действия, которые могут повлиять на работу системы
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-4">
          <div className="p-4 rounded-lg bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800">
            <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
              <div>
                <h4 className="font-medium text-red-800 dark:text-red-300">Очистка всех данных</h4>
                <p className="text-sm text-red-700 dark:text-red-400 mt-1">
                  Удаление всех пользователей и данных системы. Это действие необратимо.
                </p>
              </div>
              <Button variant="destructive" disabled>
                Очистить все данные
              </Button>
            </div>
          </div>

          <div className="p-4 rounded-lg bg-amber-50 dark:bg-amber-900/20 border border-amber-200 dark:border-amber-800">
            <div className="flex flex-col sm:flex-row justify-between items-start sm:items-center gap-4">
              <div>
                <h4 className="font-medium text-amber-800 dark:text-amber-300">Экспорт данных</h4>
                <p className="text-sm text-amber-700 dark:text-amber-400 mt-1">
                  Создание резервной копии всех данных системы.
                </p>
              </div>
              <Button variant="outline" className="border-amber-300 text-amber-700 hover:bg-amber-50">
                Экспортировать данные
              </Button>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}