import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { UserMinus, ArrowLeft } from "lucide-react";
import { useState } from "react";

export default function DeleteUser() {
  const navigate = useNavigate();
  const [userId, setUserId] = useState("");
  const [login, setLogin] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState("");

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError("");
    
    // Simulate API call
    setTimeout(() => {
      setIsLoading(false);
      if (userId || login) {
        setSuccess(true);
        setTimeout(() => {
          setSuccess(false);
          setUserId("");
          setLogin("");
        }, 3000);
      } else {
        setError("Введите ID или логин пользователя");
      }
    }, 1000);
  };

  return (
    <div className="container mx-auto px-4 py-8 max-w-2xl">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Удаление пользователя</h1>
          <p className="text-muted-foreground mt-2">
            Удаление пользователя из системы ITI
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
            <UserMinus className="h-5 w-5" />
            Удаление пользователя
          </CardTitle>
          <CardDescription>
            Удалите пользователя по ID или логину
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
                <span className="font-medium">Пользователь успешно удалён!</span>
              </div>
              <p className="text-sm text-green-600 dark:text-green-400 mt-2">
                Пользователь был удалён из системы. Это действие необратимо.
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
                  className="flex-1 bg-red-600 hover:bg-red-700"
                >
                  {isLoading ? (
                    <>
                      <svg className="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" />
                        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                      </svg>
                      Удаление...
                    </>
                  ) : (
                    "Удалить пользователя"
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

              <div className="p-4 rounded-lg bg-amber-50 dark:bg-amber-900/20 border border-amber-200 dark:border-amber-800">
                <div className="flex items-start gap-3">
                  <div className="h-5 w-5 rounded-full bg-amber-100 dark:bg-amber-800 flex items-center justify-center flex-shrink-0 mt-0.5">
                    <svg className="h-3 w-3 text-amber-600 dark:text-amber-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.998-.833-2.732 0L4.732 16.5c-.77.833.192 2.5 1.732 2.5z" />
                    </svg>
                  </div>
                  <div>
                    <p className="text-sm font-medium text-amber-800 dark:text-amber-300">Внимание!</p>
                    <p className="text-sm text-amber-700 dark:text-amber-400 mt-1">
                      Удаление пользователя — необратимое действие. Все данные пользователя будут удалены из системы.
                    </p>
                  </div>
                </div>
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
            Backend endpoint для удаления пользователя
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="p-3 rounded-lg bg-muted">
            <code className="text-sm font-mono">POST /api/private/user-manipulation/delete</code>
            <p className="text-sm text-muted-foreground mt-1">
              Требует параметры: <code className="text-xs">id</code> или <code className="text-xs">login</code>
            </p>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}