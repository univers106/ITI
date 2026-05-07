import { useState } from "react";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Field } from "@/components/ui/field";
import { Label } from "@/components/ui/label";
import { useNavigate } from "react-router-dom";
import { ArrowLeft, UserPlus, CheckCircle2 } from "lucide-react";

export default function CreateUser() {
  const navigate = useNavigate();
  const [formData, setFormData] = useState({
    login: "",
    name: "",
    password: "",
    confirmPassword: ""
  });
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData(prev => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    
    if (formData.password !== formData.confirmPassword) {
      alert("Пароли не совпадают!");
      return;
    }

    if (!formData.login || !formData.name || !formData.password) {
      alert("Все поля обязательны для заполнения!");
      return;
    }

    setLoading(true);
    
    // Simulate API call
    setTimeout(() => {
      console.log("Creating user with data:", formData);
      console.log("API endpoint: POST /api/private/user-manipulation/create");
      setLoading(false);
      setSuccess(true);
      
      // Reset form after success
      setTimeout(() => {
        setFormData({ login: "", name: "", password: "", confirmPassword: "" });
        setSuccess(false);
      }, 2000);
    }, 1000);
  };

  return (
    <div className="container mx-auto px-4 py-8 max-w-2xl">
      {/* Header */}
      <div className="flex items-center gap-4 mb-8">
        <Button variant="outline" size="icon" onClick={() => navigate("/admin")}>
          <ArrowLeft className="h-4 w-4" />
        </Button>
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Создание пользователя</h1>
          <p className="text-muted-foreground mt-2">
            Добавление нового пользователя в систему ITI
          </p>
        </div>
      </div>

      <Card>
        <CardHeader>
          <CardTitle className="flex items-center gap-2">
            <UserPlus className="h-5 w-5" />
            Форма создания пользователя
          </CardTitle>
          <CardDescription>
            Заполните все поля для создания нового пользователя
          </CardDescription>
        </CardHeader>
        <CardContent>
          {success ? (
            <div className="py-8 flex flex-col items-center justify-center text-center">
              <CheckCircle2 className="h-12 w-12 text-green-500 mb-4" />
              <h3 className="text-xl font-semibold mb-2">Пользователь успешно создан!</h3>
              <p className="text-muted-foreground">
                Пользователь "{formData.login}" был добавлен в систему.
              </p>
              <Button className="mt-6" onClick={() => setSuccess(false)}>
                Создать ещё одного пользователя
              </Button>
            </div>
          ) : (
            <form onSubmit={handleSubmit} className="space-y-6">
              <div className="space-y-4">
                <div>
                  <Label htmlFor="login">Логин</Label>
                  <input
                    id="login"
                    name="login"
                    value={formData.login}
                    onChange={handleChange}
                    placeholder="Введите логин пользователя"
                    required
                    className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                  />
                  <p className="text-xs text-muted-foreground mt-1">
                    Уникальный идентификатор для входа в систему
                  </p>
                </div>

                <div>
                  <Label htmlFor="name">Имя</Label>
                  <input
                    id="name"
                    name="name"
                    value={formData.name}
                    onChange={handleChange}
                    placeholder="Введите полное имя пользователя"
                    required
                    className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                  />
                  <p className="text-xs text-muted-foreground mt-1">
                    Отображаемое имя пользователя в системе
                  </p>
                </div>

                <div>
                  <Label htmlFor="password">Пароль</Label>
                  <input
                    id="password"
                    name="password"
                    type="password"
                    value={formData.password}
                    onChange={handleChange}
                    placeholder="Введите пароль"
                    required
                    className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                  />
                  <p className="text-xs text-muted-foreground mt-1">
                    Минимум 6 символов
                  </p>
                </div>

                <div>
                  <Label htmlFor="confirmPassword">Подтверждение пароля</Label>
                  <input
                    id="confirmPassword"
                    name="confirmPassword"
                    type="password"
                    value={formData.confirmPassword}
                    onChange={handleChange}
                    placeholder="Повторите пароль"
                    required
                    className="flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                  />
                </div>
              </div>

              <div className="p-4 rounded-lg bg-muted">
                <h4 className="font-medium mb-2">Информация о API:</h4>
                <code className="text-sm font-mono block mb-1">
                  POST /api/private/user-manipulation/create
                </code>
                <p className="text-xs text-muted-foreground">
                  Параметры: login (string), name (string), password (string)
                </p>
              </div>

              <div className="flex gap-3">
                <Button type="submit" disabled={loading} className="flex-1">
                  {loading ? "Создание..." : "Создать пользователя"}
                </Button>
                <Button type="button" variant="outline" onClick={() => navigate("/admin")}>
                  Отмена
                </Button>
              </div>
            </form>
          )}
        </CardContent>
      </Card>

      {/* Footer */}
      <div className="mt-8 text-center text-sm text-muted-foreground">
        <p>После создания пользователь сможет войти в систему с указанными данными.</p>
      </div>
    </div>
  );
}