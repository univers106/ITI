import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { useNavigate } from "react-router-dom";
import { Users, UserPlus, UserMinus, Key, Shield, Settings, ArrowLeft } from "lucide-react";

export default function AdminMain() {
  const navigate = useNavigate();

  return (
    <div className="container mx-auto px-4 py-8 max-w-6xl">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-8">
        <div>
          <h1 className="text-3xl font-bold tracking-tight">Администрирование</h1>
          <p className="text-muted-foreground mt-2">
            Управление пользователями и системой ITI
          </p>
        </div>
        <Button variant="outline" onClick={() => navigate("/dashboard")} className="gap-2">
          <ArrowLeft className="h-4 w-4" />
          Назад в кабинет
        </Button>
      </div>

      {/* Administration Cards */}
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mb-8">
        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <UserPlus className="h-5 w-5" />
              Создание пользователя
            </CardTitle>
            <CardDescription>
              Добавление нового пользователя в систему
            </CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground mb-4">
              Создайте нового пользователя с указанием логина, имени и пароля.
            </p>
            <Button className="w-full" onClick={() => navigate("/admin/create-user")}>
              Перейти к созданию
            </Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <UserMinus className="h-5 w-5" />
              Удаление пользователя
            </CardTitle>
            <CardDescription>
              Удаление пользователя из системы
            </CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground mb-4">
              Удалите пользователя по ID или логину.
            </p>
            <Button className="w-full" onClick={() => navigate("/admin/delete-user")}>
              Перейти к удалению
            </Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Key className="h-5 w-5" />
              Смена пароля
            </CardTitle>
            <CardDescription>
              Изменение пароля пользователя
            </CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground mb-4">
              Измените пароль для любого пользователя системы.
            </p>
            <Button className="w-full" onClick={() => navigate("/admin/change-password")}>
              Изменить пароль
            </Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Users className="h-5 w-5" />
              Управление данными
            </CardTitle>
            <CardDescription>
              Изменение логина и имени пользователя
            </CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground mb-4">
              Измените логин или имя существующего пользователя.
            </p>
            <Button className="w-full" onClick={() => navigate("/admin/manage-user")}>
              Управление пользователями
            </Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Shield className="h-5 w-5" />
              Управление правами
            </CardTitle>
            <CardDescription>
              Назначение прав доступа пользователям
            </CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground mb-4">
              Добавьте или удалите права доступа для пользователей.
            </p>
            <Button className="w-full" onClick={() => navigate("/admin/permissions")}>
              Управление правами
            </Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle className="flex items-center gap-2">
              <Settings className="h-5 w-5" />
              Настройки системы
            </CardTitle>
            <CardDescription>
              Общие настройки системы ITI
            </CardDescription>
          </CardHeader>
          <CardContent>
            <p className="text-sm text-muted-foreground mb-4">
              Настройки конфигурации и параметров системы.
            </p>
            <Button className="w-full" onClick={() => navigate("/admin/settings")}>
              Настройки системы
            </Button>
          </CardContent>
        </Card>
      </div>

      {/* API Information */}
      <Card>
        <CardHeader>
          <CardTitle>API Endpoints для администрирования</CardTitle>
          <CardDescription>
            Доступные endpoints в бэкенде для управления пользователями
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="space-y-3">
            <div className="p-3 rounded-lg bg-muted">
              <code className="text-sm font-mono">POST /api/private/user-manipulation/create</code>
              <p className="text-sm text-muted-foreground mt-1">Создание нового пользователя</p>
            </div>
            <div className="p-3 rounded-lg bg-muted">
              <code className="text-sm font-mono">POST /api/private/user-manipulation/delete</code>
              <p className="text-sm text-muted-foreground mt-1">Удаление пользователя</p>
            </div>
            <div className="p-3 rounded-lg bg-muted">
              <code className="text-sm font-mono">POST /api/private/user-manipulation/change-password</code>
              <p className="text-sm text-muted-foreground mt-1">Изменение пароля пользователя</p>
            </div>
            <div className="p-3 rounded-lg bg-muted">
              <code className="text-sm font-mono">POST /api/private/user-manipulation/change-login</code>
              <p className="text-sm text-muted-foreground mt-1">Изменение логина пользователя</p>
            </div>
            <div className="p-3 rounded-lg bg-muted">
              <code className="text-sm font-mono">POST /api/private/user-manipulation/change-name</code>
              <p className="text-sm text-muted-foreground mt-1">Изменение имени пользователя</p>
            </div>
          </div>
        </CardContent>
      </Card>

      {/* Footer */}
      <div className="mt-8 text-center text-sm text-muted-foreground">
        <p>Административная панель ITI • Только для авторизованных администраторов</p>
      </div>
    </div>
  );
}