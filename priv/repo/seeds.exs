alias ChinguCentral.Accounts.User
alias ChinguCentral.Applications.Cohort
alias ChinguCentral.Forums.{Category, Post, Comment}
alias ChinguCentral.Repo

# Seed users
Repo.insert!(%User{name: "Jon Doe", username: "jondoe", email: "jon@doe.com", encrypted_password: "12345"})
Repo.insert!(%User{name: "Jana Doe", username: "janadoe", email: "jana@doe.com", encrypted_password: "12345"})

# Seed forum categories
for _ <- 1..5 do
  Repo.insert!(%Category{
    name: Faker.Lorem.sentence,
    description: Faker.Lorem.paragraph
  })
end

# Seed forum posts
for _ <- 1..20 do
  Repo.insert!(%Post{
    title: Faker.Lorem.sentence,
    content: Faker.Lorem.paragraph,
    accounts_users_id: Enum.random(1..2),
    forums_categories_id: Enum.random(1..5)
  })
end

# Seed forum comments
for _ <- 1..20 do
  Repo.insert!(%Comment{
    content: Faker.Lorem.sentence,
    accounts_users_id: Enum.random(1..2),
    forums_posts_id: Enum.random(1..20)
  })
end