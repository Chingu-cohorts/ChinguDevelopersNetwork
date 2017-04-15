alias ChinguCentral.Accounts.User
alias ChinguCentral.Applications.{Cohort, UserCohort}
alias ChinguCentral.Forums.{Category, Post, Comment}
alias ChinguCentral.Repo

# Seed cohorts
Repo.insert!(%Cohort{name: "Penguin", description: "Penguins cohort"})
Repo.insert!(%Cohort{name: "Red Panda", description: "Red Pandas cohort"})
Repo.insert!(%Cohort{name: "Flamingo", description: "Flamingos cohort"})
Repo.insert!(%Cohort{name: "Rhino", description: "Rhinos cohort"})

# Seed users
for _ <- 1..20 do
  Repo.insert!(%User{
    name: Faker.Name.name,
    username: Faker.Name.first_name,
    email: Faker.Name.first_name <> "@" <> "gmail.com",
    encrypted_password: "123456"
  })
end

# Seed users to cohorts
for id <- 1..20 do
  Repo.insert!(%UserCohort{
    accounts_users_id: id,
    applications_cohorts_id: Enum.random(1..4)
  })
end

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