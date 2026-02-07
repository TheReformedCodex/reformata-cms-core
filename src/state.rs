use tera::Tera;

#[derive(Clone)]
pub struct AppState {
    pub tera: Tera,
}