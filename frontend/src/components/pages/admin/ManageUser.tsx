import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { Users, ArrowLeft } from "lucide-react";
import { useState } from "react";

export default function ManageUser() {
  const navigate = useNavigate();
  const [userId, setUserId] = useState("");
  const [login, setLogin] = useState("");
  const [newLogin, setNewLogin] = useState("");
  const [newName, setNewName] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState("");
  const [actionType, setActionType] = useState<"login" | "name" | "both">("both");

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
    
    if (actionType === "login" && !newLogin) {
      setError("Введите новый логин");
      setIsLoading(false);
      return;
    }
    
    if (actionType === "name" && !newName) {
      setError("Введите новое имя");
      setIsLoading(false);
      return;
    }
    
    if (actionType === "both" && (!newLogin || !newName)) {
      setError("Введите новый логин и имя");
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
        setNewLogin("");
        setNewName("");
      }, 3000);
    }, 1000);
  };

  return (
    <div className="container mx-auto px-4 py-8 max-w-2xl">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Управление данными пользователя</h1>
          <p className="text-muted-foreground mt-2">
            Изменение логина и имени пользователя системы ITI
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
            <Users className="h-5 w-5" />
            Изменение данных пользователя
          </CardTitle>
          <CardDescription>
            Измените логин или имя существующего пользователя
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
                <span className="font-medium">Данные пользователя успешно изменены!</span>
              </div>
              <p className="text-sm text-green-600 dark:text-green-400 mt-2">
                Данные пользователя были успешно обновлены.
              </p>
            </div>
          ) : (
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="space-y-4">
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

                <div className="relative">
                  <div className="absolute inset-0 flex items-center">
                    <span className="w-full border-t" />
                  </div>
                  <div className="relative flex justify-center text-xs uppercase">
                    <span className="bg-background px-2 text-muted-foreground">или</span>
                  </div>
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

                <div className="space-y-3">
                  <div className="flex items-center space-x-4">
                    <label className="flex items-center space-x-2">
                      <input
                        type="radio"
                        name="actionType"
                        checked={actionType === "both"}
                        onChange={() => setActionType("both")}
                        className="h-4 w-4"
                      />
                      <span className="text-sm">Изменить логин и имя</span>
                    </label>
                    <label className="flex items-center space-x-2">
                      <input
                        type="radio"
                        name="actionType"
                        checked={actionType === "login"}
                        onChange={() => setActionType("login")}
                        className="h-4 w-4"
                      />
                      <span className="text-sm">Изменить только логин</span>
                    </label>
                    <label className="flex items-center space-x-2">
                      <input
                        type="radio"
                        name="actionType"
                        checked={actionType === "name"}
                        onChange={() => setActionType("name")}
                        className="h-4 w-4"
                      />
                      <span className="text-sm">Изменить только имя</span>
                    </label>
                  </div>
                </div>

                {(actionType === "both" || actionType === "login") && (
                  <div>
                    <label htmlFor="newLogin" className="block text-sm font-medium mb-2">
                      Новый логин
                    </label>
                    <input
                      id="newLogin"
                      type="text"
                      value={newLogin}
                      onChange={(e) => setNewLogin(e.target.value)}
                      className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                      placeholder="Введите новый логин"
                    />
                  </div>
                )}

                {(actionType === "both" || actionType === "name") && (
                  <div>
                    <label htmlFor="newName" className="block text-sm font-medium mb-2">
                      Новое имя
                    </label>
                    <input
                      id="newName"
                      type="text"
                      value={newName}
                      onChange={(e) => setNewName(e.target.value)}
                      className="w-full px-3 py-2 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                      placeholder="Введите новое имя"
                    />
                  </div>
                )}
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
                      Изменение...
                    </>
                  ) : (
                    "Изменить данные"
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
          <CardTitle>API Endpoints</CardTitle>
          <CardDescription>
            Backend endpoints для изменения данных пользователя
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            <div className="p-3 rounded-lg bg-muted">
              <code className="text-sm font-mono">POST /api/private/user-manipulation/change-login</code>
              <p className="text-sm text-muted-foreground mt-1">
                Требует параметры: <code className="text-xs">id</code> или <code className="text-xs">login</code>, <code className="text-xs">new_login</code>
              </p>
            </div>
            <div className="p-3 rounded-lg bg-muted">
              <code className="text-sm font-mono">POST /api/private/user-manipulation/change-name</code>
              <p className="text-sm text-muted-foreground mt-1">
                Требует параметры: <code className="text-xs">id</code> или <code className="text-xs">login</code>, <code className="text-xs">new_name</code>
              </p>
            </div>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}