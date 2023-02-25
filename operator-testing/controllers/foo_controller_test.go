package controllers

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	tutorialv1 "my.domain/tutorial/api/v1"
)

var _ = Describe("Controllers/FooController", func() {

	const (
		foo1Name = "foo-1"

		//foo2Name = "foo-2"

		namespace = "default"
	)

	ctx = context.Background()
	By("Creating a first Foo custom resource")

	foo1 := &tutorialv1.Foo{
		ObjectMeta: metav1.ObjectMeta{
			Name:      foo1Name,
			Namespace: namespace,
		},
		Spec: tutorialv1.FooSpec{
			Name: foo1Name,
		},
	}

	Context("When setting up test env", func() {

		It("Should be able to create first custrom resource", func() {

			Expect(k8sClient.Create(ctx, foo1)).Should(Succeed())

		})

		It("Should be able to update the resource", func() {

			foo1.Spec.Name = "shubham"

			Expect(k8sClient.Update(ctx, foo1, &client.UpdateOptions{})).Should(Succeed())

		})

		It("Should be able to create 2nd custom resource", func() {

			By("Creating the second custorm resource")

			Expect(k8sClient.Delete(ctx, foo1, &client.DeleteAllOfOptions{})).Should(Succeed())

		})

	})

})
