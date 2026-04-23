import { Button } from "@/components/ui/button";
import { Link } from "react-router-dom";

export default function NotFoundPage() {
  return (
    <main className="min-h-screen flex items-center justify-center bg-background px-4 py-8">
      <div className="max-w-xl w-full text-center space-y-6">
        <p className="text-sm font-medium uppercase tracking-widest text-muted-foreground">
          404
        </p>
        <h1 className="text-5xl font-bold tracking-tight">
          Страница не найдена
        </h1>
        <p className="text-base text-muted-foreground">
          К сожалению, мы не смогли найти страницу, которую вы ищете. Возможно,
          URL введен неверно или страница была удалена.
        </p>

        <div className="flex justify-center gap-3">
          <Button size="lg">
            <Link to="/">На главную</Link>
          </Button>
        </div>
      </div>
    </main>
  );
}
