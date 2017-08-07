import { mount } from 'avoriaz'
import Home from '@/components/Home'

describe('Home.vue', () => {
  it('should render home component', () => {
    const wrapper = mount(Home)
    expect(wrapper.is('div')).to.equal(true)
    expect(wrapper.hasClass('home')).to.equal(true)
  })

  it('should display the correct title', () => {
    const expectedMessage = 'Cohorts'
    const wrapper = mount(Home)
    expect(wrapper.find('.title')[0].text()).to.equal(expectedMessage)
  })
})
