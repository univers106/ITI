import { useState } from "react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Field } from "@/components/ui/field";
import { Label } from "@/components/ui/label";
import { login, type ApiError } from "@/lib/api";
import { useNavigate } from "react-router-dom";
import { AlertCircle, CheckCircle2 } from "lucide-react";

export default function AuthPage() {
  const navigate = useNavigate();
  const [loginValue, setLogin] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState(false);
  const [loading, setLoading] = useState(false);
  const [fieldErrors, setFieldErrors] = useState<{ login?: string; password?: string }>({});

  const validateForm = () => {
    const errors: { login?: string; password?: string } = {};
    if (!loginValue.trim()) {
      errors.login = "Логин обязателен";
    }
    if (!password) {
      errors.password = "Пароль обязателен";
    } else if (password.length < 3) {
      errors.password = "Пароль должен быть не менее 3 символов";
    }
    setFieldErrors(errors);
    return Object.keys(errors).length === 0;
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);
    setSuccess(false);

    if (!validateForm()) {
      return;
    }

    setLoading(true);

    try {
      await login({ login: loginValue, password });
      setSuccess(true);
      // Успешный вход - перезагружаем страницу для обновления состояния
      // Используем небольшую задержку чтобы пользователь увидел сообщение
      setTimeout(() => {
        // Принудительная перезагрузка с переходом на главную
        // Добавляем timestamp чтобы избежать кэширования
        const timestamp = new Date().getTime();
        window.location.href = `/?_=${timestamp}`;
      }, 800);
    } catch (err) {
      const apiError = err as ApiError;
      setError(apiError.message || "Login failed. Please check your credentials.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container mx-auto px-4 py-16 max-w-md">
      <Card>
        <CardHeader>
          <CardTitle className="text-2xl">Вход в систему</CardTitle>
          <CardDescription>
            Введите ваш логин и пароль для доступа к системе
          </CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={handleSubmit} className="space-y-6">
            <Field>
              <Label htmlFor="login">Логин</Label>
              <input
                id="login"
                type="text"
                value={loginValue}
                onChange={(e) => {
                  setLogin(e.target.value);
                  if (fieldErrors.login) setFieldErrors({ ...fieldErrors, login: undefined });
                }}
                className={`flex h-10 w-full rounded-md border ${fieldErrors.login ? "border-destructive" : "border-input"} bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50`}
                placeholder="Введите логин"
                disabled={loading}
              />
              {fieldErrors.login && (
                <p className="text-sm text-destructive mt-1 flex items-center gap-1">
                  <AlertCircle className="h-3 w-3" />
                  {fieldErrors.login}
                </p>
              )}
            </Field>
            <Field>
              <Label htmlFor="password">Пароль</Label>
              <input
                id="password"
                type="password"
                value={password}
                onChange={(e) => {
                  setPassword(e.target.value);
                  if (fieldErrors.password) setFieldErrors({ ...fieldErrors, password: undefined });
                }}
                className={`flex h-10 w-full rounded-md border ${fieldErrors.password ? "border-destructive" : "border-input"} bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50`}
                placeholder="Введите пароль"
                disabled={loading}
              />
              {fieldErrors.password && (
                <p className="text-sm text-destructive mt-1 flex items-center gap-1">
                  <AlertCircle className="h-3 w-3" />
                  {fieldErrors.password}
                </p>
              )}
            </Field>
            {error && (
              <div className="text-sm text-destructive bg-destructive/10 p-3 rounded-md flex items-center gap-2">
                <AlertCircle className="h-4 w-4" />
                {error}
              </div>
            )}
            {success && (
              <div className="text-sm text-green-600 bg-green-50 p-3 rounded-md flex items-center gap-2">
                <CheckCircle2 className="h-4 w-4" />
                Успешный вход! Перенаправление...
              </div>
            )}
            <Button type="submit" className="w-full" disabled={loading}>
              {loading ? "Вход..." : "Войти"}
            </Button>
          </form>
          <div className="mt-6 text-center text-sm text-muted-foreground">
            <p>
              Для тестирования используйте логин "test_user" и пароль "test_password"
            </p>
            <p className="mt-2">
              Или логин "test_admin" с тем же паролем для администратора.
            </p>
          </div>
        </CardContent>
      </Card>
    </div>
  );
}