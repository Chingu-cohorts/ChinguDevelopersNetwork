import { mount } from 'avoriaz'
import Navbar from '@/components/misc/Navbar'

describe('Navbar.vue', () => {
  it('should render the navbar component', () => {
    const wrapper = mount(Navbar)
    expect(wrapper.is('nav')).to.equal(true)
  })

  it('should have a computed property to get the logged user', () => {
    const wrapper = mount(Navbar)
    expect(typeof wrapper.computed().loggedUser).to.equal('function')
  })

  it('should have a computed property to get the user gravatar', () => {
    const wrapper = mount(Navbar)
    expect(typeof wrapper.computed().userGravatar).to.equal('function')
  })

  it('should have a collapse navbar method', () => {
    const wrapper = mount(Navbar)
    expect(typeof wrapper.methods().collapse).to.equal('function')
  })

  it('should have a reload user method', () => {
    const wrapper = mount(Navbar)
    expect(typeof wrapper.methods().reloadUser).to.equal('function')
  })

  it('should have a logout method', () => {
    const wrapper = mount(Navbar)
    expect(typeof wrapper.methods().logout).to.equal('function')
  })
})
