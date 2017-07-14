import md5 from 'blueimp-md5'

export const gravatar = (email) => {
  let hash = md5(email)
  let gravatarUrl = `https://gravatar.com/avatar/${hash}?s=512`
  return gravatarUrl
}

export const reputation = (experience) => {
  let reputation

  let assignRep = (name) => {
    reputation = name
  }

  switch (true) {
    case experience >= 1000:
      assignRep('Legend')
      break
    case experience >= 500:
      assignRep('Sage')
      break
    case experience >= 100:
      assignRep('Senior')
      break
    case experience >= 10:
      assignRep('Member')
      break
    default:
      assignRep('Neutral')
      break
  }

  return reputation
}
